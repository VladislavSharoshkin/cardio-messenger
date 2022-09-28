import 'dart:async';

import 'package:cardio_messenger/models/user.dart';
import 'package:cardio_messenger/services/get_participants.dart';
import 'package:cardio_messenger/services/get_users.dart';
import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/l10n.dart';
import 'package:future_loading_dialog/future_loading_dialog.dart';
import 'package:matrix/matrix.dart';
import 'package:vrouter/vrouter.dart';

import 'package:cardio_messenger/pages/invitation_selection/invitation_selection_view.dart';
import 'package:cardio_messenger/widgets/matrix.dart';
import '../../utils/localized_exception_extension.dart';

class InvitationSelection extends StatefulWidget {
  const InvitationSelection({Key? key}) : super(key: key);

  @override
  InvitationSelectionController createState() =>
      InvitationSelectionController();
}

class InvitationSelectionController extends State<InvitationSelection> {
  TextEditingController controller = TextEditingController();
  late String currentSearchTerm;
  bool loading = false;
  List<MUser> foundProfiles = [];
  Timer? coolDown;

  String? get roomId => VRouter.of(context).pathParameters['roomid'];

  // Future<List<User>> getContacts(BuildContext context) async {
  //   final client = Matrix.of(context).client;
  //   final room = client.getRoomById(roomId!)!;
  //   final participants = await room.requestParticipants();
  //   participants.removeWhere(
  //     (u) => ![Membership.join, Membership.invite].contains(u.membership),
  //   );
  //   final participantsIds = participants.map((p) => p.stateKey).toList();
  //   final contacts = client.rooms
  //       .where((r) => r.isDirectChat)
  //       .map((r) => r.getUserByMXIDSync(r.directChatMatrixID!))
  //       .toList()
  //     ..removeWhere((u) => participantsIds.contains(u.stateKey));
  //   contacts.sort(
  //     (a, b) => a.calcDisplayname().toLowerCase().compareTo(
  //           b.calcDisplayname().toLowerCase(),
  //         ),
  //   );
  //   return contacts;
  // }

  void inviteAction(BuildContext context, int id) async {
    final room = Matrix.of(context).client.getRoomById(roomId!);
    final success = await showFutureLoadingDialog(
      context: context,
      future: () => GetParticipants.invite(int.parse(roomId!), id),//room!.invite(id),
    );
    if (success.error == null) {
      ScaffoldMessenger.of(context).showSnackBar(SnackBar(
          content: Text(L10n.of(context)!.contactHasBeenInvitedToTheGroup)));
    }
  }

  void searchUserWithCoolDown(String text) async {
    coolDown?.cancel();
    coolDown = Timer(
      const Duration(seconds: 1),
      () => searchUser(context, text),
    );
  }

  void searchUser(BuildContext context, String text) async {
    coolDown?.cancel();
    if (text.isEmpty) {
      setState(() => foundProfiles = []);
    }
    currentSearchTerm = text;
    if (currentSearchTerm.isEmpty) return;
    if (loading) return;
    setState(() => loading = true);

    GetUsers response;
    try {
      response = await GetUsers.search(text);
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text((e).toLocalizedString(context))));
      return;
    } finally {
      setState(() => loading = false);
    }
    setState(() {
      foundProfiles = response.users;
      // if (text.isValidMatrixId &&
      //     foundProfiles.indexWhere((profile) => text == profile.userId) == -1) {
      //   setState(() => foundProfiles = [
      //         Profile.fromJson({'user_id': text}),
      //       ]);
      // }
      // final participants = Matrix.of(context)
      //     .client
      //     .getRoomById(roomId!)!
      //     .getParticipants()
      //     .where((user) =>
      //         [Membership.join, Membership.invite].contains(user.membership))
      //     .toList();
      // foundProfiles.removeWhere((profile) =>
      //     participants.indexWhere((u) => u.id == profile.userId) != -1);
    });
  }

  @override
  Widget build(BuildContext context) => InvitationSelectionView(this);
}
