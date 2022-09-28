import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/participant.dart';
import 'package:cardio_messenger/models/user.dart';
import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/l10n.dart';
import 'package:matrix/matrix.dart';
import 'package:cardio_messenger/utils/int_uri.dart';
import '../../widgets/avatar.dart';
import '../user_bottom_sheet/user_bottom_sheet.dart';

class ParticipantListItem extends StatelessWidget {
  final MParticipant user;
  final MChat chat;

  const ParticipantListItem(this.user, this.chat, {Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final membershipBatch = <Membership, String>{
      Membership.join: '',
      Membership.ban: L10n.of(context)!.banned,
      Membership.invite: L10n.of(context)!.invited,
      Membership.leave: L10n.of(context)!.leftTheChat,
    };
    // final permissionBatch = user.powerLevel == 100
    //     ? L10n.of(context)!.admin
    //     : user.powerLevel >= 50
    //         ? L10n.of(context)!.moderator
    //         : '';
    final permissionBatch = user.user.id == chat.creatorId ? L10n.of(context)!.admin : '';

    return Opacity(
      // opacity: user.membership == Membership.join ? 1 : 0.5,
      opacity: 1,
      child: ListTile(
        onTap: () => showModalBottomSheet(
          context: context,
          builder: (c) => UserBottomSheet(
            user: user.user,
            outerContext: context,
          ),
        ),
        title: Row(
          children: <Widget>[
            Text(user.user.getFullName()),
            permissionBatch.isEmpty
                ? Container()
                : Container(
                    padding: const EdgeInsets.all(4),
                    margin: const EdgeInsets.symmetric(horizontal: 8),
                    decoration: BoxDecoration(
                      color: Theme.of(context).secondaryHeaderColor,
                      borderRadius: BorderRadius.circular(8),
                    ),
                    child: Center(child: Text(permissionBatch)),
                  ),
            membershipBatch[Membership.join]!.isEmpty
                ? Container()
                : Container(
                    padding: const EdgeInsets.all(4),
                    margin: const EdgeInsets.symmetric(horizontal: 8),
                    decoration: BoxDecoration(
                      color: Theme.of(context).secondaryHeaderColor,
                      borderRadius: BorderRadius.circular(8),
                    ),
                    child:
                        Center(child: Text(membershipBatch[Membership.join]!)),
                  ),
          ],
        ),
        subtitle: Text(user.user.username),
        leading:
            Avatar(mxContent: user.user.avatarID?.uri, name: user.user.getFullName()),
      ),
    );
  }
}
