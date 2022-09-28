import 'dart:async';
import 'dart:io';
import 'package:cardio_messenger/push_notification.dart';
import 'package:cardio_messenger/socket_helper.dart';
import 'package:cardio_messenger/store_helper.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:adaptive_theme/adaptive_theme.dart';
import 'package:flutter_app_lock/flutter_app_lock.dart';
import 'package:flutter_gen/gen_l10n/l10n.dart';
import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:future_loading_dialog/future_loading_dialog.dart';
//import 'package:gaeilge_flutter_l10n/gaeilge_flutter_l10n.dart';
import 'package:matrix/matrix.dart';
import 'package:universal_html/html.dart' as html;
import 'package:vrouter/vrouter.dart';
import 'package:cardio_messenger/config/routes.dart';
import 'package:cardio_messenger/utils/client_manager.dart';
import 'package:cardio_messenger/utils/platform_infos.dart';
import 'package:cardio_messenger/utils/sentry_controller.dart';
import 'config/app_config.dart';
import 'config/themes.dart';
import 'utils/background_push.dart';
import 'utils/custom_scroll_behaviour.dart';
import 'utils/localized_exception_extension.dart';
import 'utils/platform_infos.dart';
import 'widgets/lock_screen.dart';
import 'widgets/matrix.dart';

class MyHttpOverrides extends HttpOverrides{
  @override
  HttpClient createHttpClient(SecurityContext? context){
    return super.createHttpClient(context)
      ..badCertificateCallback = (X509Certificate cert, String host, int port)=> true;
  }
}

void main() async {
  // Our background push shared isolate accesses flutter-internal things very early in the startup proccess
  // To make sure that the parts of flutter needed are started up already, we need to ensure that the
  // widget bindings are initialized already.
  HttpOverrides.global = MyHttpOverrides();
  WidgetsFlutterBinding.ensureInitialized();
  await StoreHelper.shared.init();

  FlutterError.onError =
      (FlutterErrorDetails details) => Zone.current.handleUncaughtError(
            details.exception,
            details.stack ?? StackTrace.current,
          );

  final clients = await ClientManager.getClients();
  Logs().level = kReleaseMode ? Level.warning : Level.verbose;

  if (PlatformInfos.isMobile) {
    //BackgroundPush.clientOnly(clients.first);
    PushNotification.shared.init();
  }

  final queryParameters = <String, String>{};
  if (kIsWeb) {
    queryParameters
        .addAll(Uri.parse(html.window.location.href).queryParameters);
  }

  runZonedGuarded(
    () => runApp(PlatformInfos.isMobile
        ? AppLock(
            builder: (args) => CardioMessengerApp(
              clients: clients,
              queryParameters: queryParameters,
            ),
            lockScreen: const LockScreen(),
            enabled: false,
          )
        : CardioMessengerApp(clients: clients, queryParameters: queryParameters)),
    SentryController.captureException,
  );
}

class CardioMessengerApp extends StatefulWidget {
  final Widget? testWidget;
  final List<Client> clients;
  final Map<String, String>? queryParameters;

  const CardioMessengerApp({
    Key? key,
    this.testWidget,
    required this.clients,
    this.queryParameters,
  }) : super(key: key);

  /// getInitialLink may rereturn the value multiple times if this view is
  /// opened multiple times for example if the user logs out after they logged
  /// in with qr code or magic link.
  static bool gotInitialLink = false;

  @override
  _CardioMessengerAppState createState() => _CardioMessengerAppState();
}

class _CardioMessengerAppState extends State<CardioMessengerApp> {
  GlobalKey<VRouterState>? _router;
  bool? columnMode;
  String? _initialUrl;

  @override
  void initState() {
    super.initState();
    // _initialUrl =
    //     widget.clients.any((client) => client.isLogged()) ? '/rooms' : '/home';
    _initialUrl = StoreHelper.shared.isLogged() ? '/rooms' : '/home';
    if (StoreHelper.shared.isLogged()) {
      SocketHelper.shared.connectSocket();
    }
  }

  @override
  Widget build(BuildContext context) {
    return AdaptiveTheme(
      light: FluffyThemes.light,
      dark: FluffyThemes.dark,
      initial: AdaptiveThemeMode.system,
      builder: (theme, darkTheme) => LayoutBuilder(
        builder: (context, constraints) {
          const maxColumns = 3;
          var newColumns =
              (constraints.maxWidth / FluffyThemes.columnWidth).floor();
          if (newColumns > maxColumns) newColumns = maxColumns;
          columnMode ??= newColumns > 1;
          _router ??= GlobalKey<VRouterState>();
          if (columnMode != newColumns > 1) {
            Logs().v('Set Column Mode = $columnMode');
            WidgetsBinding.instance?.addPostFrameCallback((_) {
              setState(() {
                _initialUrl = _router?.currentState?.url;
                columnMode = newColumns > 1;
                _router = GlobalKey<VRouterState>();
              });
            });
          }
          return VRouter(
            key: _router,
            title: AppConfig.applicationName,
            theme: theme,
            scrollBehavior: CustomScrollBehavior(),
            logs: kReleaseMode ? VLogs.none : VLogs.info,
            darkTheme: darkTheme,
            localizationsDelegates: const [
              ...L10n.localizationsDelegates,
              //GaMaterialLocalizations.delegate
            ],
            supportedLocales: L10n.supportedLocales,
            initialUrl: _initialUrl ?? '/',
            routes: AppRoutes(columnMode ?? false).routes,
            builder: (context, child) {
              LoadingDialog.defaultTitle = L10n.of(context)!.loadingPleaseWait;
              LoadingDialog.defaultBackLabel = L10n.of(context)!.close;
              LoadingDialog.defaultOnError =
                  (e) => (e as Object?)!.toLocalizedString(context);
              WidgetsBinding.instance?.addPostFrameCallback((_) {
                SystemChrome.setSystemUIOverlayStyle(
                  SystemUiOverlayStyle(
                    statusBarColor: Colors.transparent,
                    systemNavigationBarColor:
                        Theme.of(context).appBarTheme.backgroundColor,
                    systemNavigationBarIconBrightness:
                        Theme.of(context).brightness == Brightness.light
                            ? Brightness.dark
                            : Brightness.light,
                  ),
                );
              });
              return Matrix(
                context: context,
                router: _router,
                clients: widget.clients,
                child: child,
              );
            },
          );
        },
      ),
    );
  }
}
