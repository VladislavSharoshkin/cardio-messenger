import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/services/get_user.dart';
import 'package:cardio_messenger/utils/date_time_extension.dart';
import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/l10n.dart';
import 'package:vrouter/vrouter.dart';
import 'package:cardio_messenger/utils/int_uri.dart';
import 'package:cardio_messenger/pages/chat/chat.dart';
import 'package:cardio_messenger/pages/user_bottom_sheet/user_bottom_sheet.dart';
import 'package:cardio_messenger/utils/matrix_sdk_extensions.dart/matrix_locals.dart';
import 'package:cardio_messenger/widgets/avatar.dart';

class ChatAppBarTitle extends StatelessWidget {
  final ChatController controller;
  const ChatAppBarTitle(this.controller, {Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final room = controller.chat;
    if (room == null) {
      return Container();
    }
    if (controller.selectedEvents.isNotEmpty) {
      return Text(controller.selectedEvents.length.toString());
    }
    final directChatMatrixID =
        room.type == ChatType.Single ? room.companion!.id : null;
    return InkWell(
      splashColor: Colors.transparent,
      highlightColor: Colors.transparent,
      onTap: directChatMatrixID != null
          ? () async {
              var getUser = await GetUser.get(directChatMatrixID);
              showModalBottomSheet(
                context: context,
                builder: (c) => UserBottomSheet(
                  user: getUser,
                  outerContext: context,
                  // onMention: () => controller.sendController.text +=
                  // '${room.getUserByMXIDSync(directChatMatrixID).mention} ',
                ),
              );
            }
          : () => VRouter.of(context)
              .toSegments(['rooms', room.id.toString(), 'details']),
      child: Row(
        children: [
          Avatar(
            mxContent: room.type == ChatType.Single
                ? room.companion!.avatarID?.uri
                : room.avatarID?.uri,
            name: room.type == ChatType.Single
                ? room.companion!.getFullName()
                : room.name,
            size: 32,
          ),
          const SizedBox(width: 12),
          Expanded(
              child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                room.type == ChatType.Single
                    ? room.companion!.getFullName()
                    : room.name,
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: const TextStyle(
                  fontSize: 16,
                ),
              ),
              Text(
                room.type == ChatType.Single
                    ? room.companion!.lastOnline != null ? room.companion!.lastOnline!.isOnline()
                        ? L10n.of(context)!.online
                        : room.companion!.lastOnline!.createdAt
                            .localizedTime(context) : 'не появлялся'
                    : '${room.participantCount} участников',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: const TextStyle(
                  fontSize: 14,
                ),
              ),
            ],
          )),
        ],
      ),
    );
  }
}
