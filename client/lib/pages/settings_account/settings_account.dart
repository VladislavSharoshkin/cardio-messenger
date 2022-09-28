import 'package:cardio_messenger/models/user.dart';
import 'package:cardio_messenger/services/get_user.dart';
import 'package:cardio_messenger/store_helper.dart';
import 'package:flutter/material.dart';

import 'package:adaptive_dialog/adaptive_dialog.dart';
import 'package:flutter_gen/gen_l10n/l10n.dart';
import 'package:future_loading_dialog/future_loading_dialog.dart';
import 'package:matrix/matrix.dart';
import 'package:vrouter/vrouter.dart';

import 'package:cardio_messenger/pages/settings_account/settings_account_view.dart';
import 'package:cardio_messenger/widgets/matrix.dart';

class SettingsAccount extends StatefulWidget {
  const SettingsAccount({Key? key}) : super(key: key);

  @override
  SettingsAccountController createState() => SettingsAccountController();
}

class SettingsAccountController extends State<SettingsAccount> {
  Future<dynamic>? profileFuture;
  MUser? profile;
  bool profileUpdated = false;

  void updateProfile() => setState(() {
        profileUpdated = true;
        profile = profileFuture = null;
      });

  void setDisplaynameAction() async {
    final input = await showTextInputDialog(
      useRootNavigator: false,
      context: context,
      title: L10n.of(context)!.editDisplayname,
      okLabel: L10n.of(context)!.ok,
      cancelLabel: L10n.of(context)!.cancel,
      textFields: [
        DialogTextField(
          initialText: profile?.getFullName(),
        )
      ],
    );
    if (input == null) return;
    final matrix = Matrix.of(context);
    final success = await showFutureLoadingDialog(
      context: context,
      future: () =>
          matrix.client.setDisplayName(matrix.client.userID!, input.single),
    );
    if (success.error == null) {
      updateProfile();
    }
  }

  void logoutAction() async {
    if (await showOkCancelAlertDialog(
          useRootNavigator: false,
          context: context,
          title: L10n.of(context)!.areYouSureYouWantToLogout,
          okLabel: L10n.of(context)!.yes,
          cancelLabel: L10n.of(context)!.cancel,
        ) ==
        OkCancelResult.cancel) {
      return;
    }
    // final matrix = Matrix.of(context);
    // await showFutureLoadingDialog(
    //   context: context,
    //   future: () => matrix.client.logout(),
    // );
    StoreHelper.shared.storage.deleteAll();
  }

  void deleteAccountAction() async {
    if (await showOkCancelAlertDialog(
          useRootNavigator: false,
          context: context,
          title: L10n.of(context)!.warning,
          message: L10n.of(context)!.deactivateAccountWarning,
          okLabel: L10n.of(context)!.ok,
          cancelLabel: L10n.of(context)!.cancel,
        ) ==
        OkCancelResult.cancel) {
      return;
    }
    if (await showOkCancelAlertDialog(
          useRootNavigator: false,
          context: context,
          title: L10n.of(context)!.areYouSure,
          okLabel: L10n.of(context)!.yes,
          cancelLabel: L10n.of(context)!.cancel,
        ) ==
        OkCancelResult.cancel) {
      return;
    }
    final input = await showTextInputDialog(
      useRootNavigator: false,
      context: context,
      title: L10n.of(context)!.pleaseEnterYourPassword,
      okLabel: L10n.of(context)!.ok,
      cancelLabel: L10n.of(context)!.cancel,
      textFields: [
        const DialogTextField(
          obscureText: true,
          hintText: '******',
          minLines: 1,
          maxLines: 1,
        )
      ],
    );
    if (input == null) return;
    await showFutureLoadingDialog(
      context: context,
      future: () => Matrix.of(context).client.deactivateAccount(
            auth: AuthenticationPassword(
              password: input.single,
              identifier: AuthenticationUserIdentifier(
                  user: Matrix.of(context).client.userID!),
            ),
          ),
    );
  }

  void addAccountAction() => VRouter.of(context).to('add');

  @override
  Widget build(BuildContext context) {
    profileFuture ??= GetUser.get(StoreHelper.shared.userID!)
        .then((p) {
      if (mounted) setState(() => profile = p);
      return p;
    });
    return SettingsAccountView(this);
  }
}
