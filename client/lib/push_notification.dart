import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:firebase_core/firebase_core.dart';
import 'dart:io' show Platform;

import 'package:flutter_local_notifications/flutter_local_notifications.dart';

Future _firebaseMessagingBackgroundHandler(RemoteMessage message) async {
  print("Handling a background message: ${message.messageId}");
}

class PushNotification {
  static final shared = PushNotification();

  late FirebaseMessaging _messaging;
  String? _token;

  String? getToken() {
    return _token;
  }

  Future<void> init() async {

    await Firebase.initializeApp();
    await FlutterLocalNotificationsPlugin().cancelAll();

    _messaging = FirebaseMessaging.instance;
    FirebaseMessaging.onBackgroundMessage(_firebaseMessagingBackgroundHandler);

    if (Platform.isIOS) {
      //_messaging.requestNotificationPermissions(IosNotificationSettings());
    }

    _token = await _messaging.getToken();
    print("FirebaseMessaging token: $_token");

    // 3. On iOS, this helps to take the user permissions
    NotificationSettings settings = await _messaging.requestPermission(
      alert: true,
      announcement: false,
      badge: true,
      carPlay: false,
      criticalAlert: false,
      provisional: false,
      sound: true,
    );

    if (settings.authorizationStatus == AuthorizationStatus.authorized) {
      print('User granted permission');
      // TODO: handle the received notifications
    } else {
      print('User declined or has not accepted permission');
    }

    FirebaseMessaging.onMessage.listen((RemoteMessage message) {
      // Parse the message received
      print("new push");
    });

    FirebaseMessaging.onMessageOpenedApp.listen((message) async {
      await FlutterLocalNotificationsPlugin().cancelAll();
    });
  }
}