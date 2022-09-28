import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/services/get_message.dart';
import 'package:cardio_messenger/store_helper.dart';
import 'package:cardio_messenger/utils/int_uri.dart';
import 'package:flutter/material.dart';

import 'package:adaptive_dialog/adaptive_dialog.dart';
import 'package:flutter_gen/gen_l10n/l10n.dart';
import 'package:future_loading_dialog/future_loading_dialog.dart';
import 'package:matrix/matrix.dart';
import 'package:pedantic/pedantic.dart';
import 'package:vrouter/vrouter.dart';

import 'package:cardio_messenger/config/app_config.dart';
import 'package:cardio_messenger/utils/matrix_sdk_extensions.dart/matrix_locals.dart';
import 'package:cardio_messenger/utils/room_status_extension.dart';
import '../../utils/date_time_extension.dart';
import '../../widgets/avatar.dart';
import '../../widgets/matrix.dart';
import '../chat/send_file_dialog.dart';

enum ArchivedRoomAction { delete, rejoin }

class ChatListItem extends StatelessWidget {
  final MChat room;
  final bool activeChat;
  final bool selected;
  final Function? onForget;
  final Function? onTap;
  final Function? onLongPress;

  const ChatListItem(
    this.room, {
    this.activeChat = false,
    this.selected = false,
    this.onTap,
    this.onLongPress,
    this.onForget,
    Key? key,
  }) : super(key: key);

  dynamic clickAction(BuildContext context) async {
    if (onTap != null) return onTap!();
    if (!activeChat) {
      // if (room.membership == Membership.invite &&
      //     (await showFutureLoadingDialog(
      //                 context: context,
      //                 future: () async {
      //                   final joinedFuture = room.client.onSync.stream
      //                       .where((u) =>
      //                           u.rooms?.join?.containsKey(room.id) ?? false)
      //                       .first;
      //                   await room.join();
      //                   await joinedFuture;
      //                 }))
      //             .error !=
      //         null) {
      //   return;
      // }
      //
      // if (room.membership == Membership.ban) {
      //   ScaffoldMessenger.of(context).showSnackBar(
      //     SnackBar(
      //       content: Text(L10n.of(context)!.youHaveBeenBannedFromThisChat),
      //     ),
      //   );
      //   return;
      // }
      //
      // if (room.membership == Membership.leave) {
      //   final action = await showModalActionSheet<ArchivedRoomAction>(
      //     context: context,
      //     title: L10n.of(context)!.archivedRoom,
      //     message: L10n.of(context)!.thisRoomHasBeenArchived,
      //     actions: [
      //       SheetAction(
      //         label: L10n.of(context)!.rejoin,
      //         key: ArchivedRoomAction.rejoin,
      //       ),
      //       SheetAction(
      //         label: L10n.of(context)!.delete,
      //         key: ArchivedRoomAction.delete,
      //         isDestructiveAction: true,
      //       ),
      //     ],
      //   );
      //   if (action != null) {
      //     switch (action) {
      //       case ArchivedRoomAction.delete:
      //         await archiveAction(context);
      //         break;
      //       case ArchivedRoomAction.rejoin:
      //         await showFutureLoadingDialog(
      //           context: context,
      //           future: () => room.join(),
      //         );
      //         break;
      //     }
      //   }
      // }

      if (true) {//room.membership == Membership.join) {
        if (Matrix.of(context).shareContent != null) {
          // if (Matrix.of(context).shareContent!['msgtype'] ==
          //     'chat.fluffy.shared_file') {
          //   await showDialog(
          //     context: context,
          //     useRootNavigator: false,
          //     builder: (c) => SendFileDialog(
          //       file: Matrix.of(context).shareContent!['file'],
          //       room: room,
          //     ),
          //   );
          // } else {
          //   unawaited(room.sendEvent(Matrix.of(context).shareContent!));
          // }
          GetMessage.send(room.id, null, Matrix.of(context).shareContent!, []);
          Matrix.of(context).shareContent = null;
        }
        VRouter.of(context).toSegments(['rooms', room.id.toString()]);
      }
      VRouter.of(context).toSegments(['rooms', room.id.toString()]);
    }
  }
  //
  // Future<void> archiveAction(BuildContext context) async {
  //   {
  //     if ([Membership.leave, Membership.ban].contains(room.membership)) {
  //       final success = await showFutureLoadingDialog(
  //         context: context,
  //         future: () => room.forget(),
  //       );
  //       if (success.error == null) {
  //         if (onForget != null) onForget!();
  //       }
  //       return;
  //     }
  //     final confirmed = await showOkCancelAlertDialog(
  //       useRootNavigator: false,
  //       context: context,
  //       title: L10n.of(context)!.areYouSure,
  //       okLabel: L10n.of(context)!.yes,
  //       cancelLabel: L10n.of(context)!.no,
  //     );
  //     if (confirmed == OkCancelResult.cancel) return;
  //     await showFutureLoadingDialog(
  //         context: context, future: () => room.leave());
  //     return;
  //   }
  // }

  @override
  Widget build(BuildContext context) {
    const isMuted = false; //room.pushRuleState != PushRuleState.notify;
    const typingText = '';//"room.getLocalizedTypingText(context)";
    final ownMessage =
        room.lastMessage?.senderId == StoreHelper.shared.userID;
    // final unread = room.isUnread || room.membership == Membership.invite;
    const unread = false;
    final unreadBubbleSize = room.unreadCount > 0
        ? 20.0
        : 14.0;
    // const unreadBubbleSize = 0.0;

    return ListTile(
      selected: selected || activeChat,
      selectedTileColor: selected
          ? Theme.of(context).primaryColor.withAlpha(100)
          : Theme.of(context).secondaryHeaderColor,
      onLongPress: onLongPress as void Function()?,
      leading: selected
          ? SizedBox(
              width: Avatar.defaultSize,
              height: Avatar.defaultSize,
              child: Material(
                color: Theme.of(context).primaryColor,
                borderRadius: BorderRadius.circular(Avatar.defaultSize),
                child: const Icon(Icons.check, color: Colors.white),
              ),
            )
          : Avatar(
              mxContent: room.type == ChatType.Single ? room.companion!.avatarID?.uri : room.avatarID?.uri,
              name: room.type == ChatType.Single ? room.companion!.getFullName() : room.name,
              onTap: onLongPress as void Function()?,
            ),
      title: Row(
        children: <Widget>[
          Expanded(
            child: Text(
              room.type == ChatType.Single ? room.companion!.getFullName() : room.name,
              maxLines: 1,
              overflow: TextOverflow.ellipsis,
              softWrap: false,
              style: TextStyle(
                fontWeight: FontWeight.bold,
                color: unread
                    ? Theme.of(context).colorScheme.secondary
                    : Theme.of(context).textTheme.bodyText1!.color,
              ),
            ),
          ),
          if (isMuted)
            const Padding(
              padding: EdgeInsets.only(left: 4.0),
              child: Icon(
                Icons.notifications_off_outlined,
                size: 16,
              ),
            ),
          if (false) // room.isFavourite
            Padding(
              padding: EdgeInsets.only(
                  right: room.unreadCount > 0 ? 4.0 : 0.0),
              child: Icon(
                Icons.push_pin_outlined,
                size: 16,
                color: Theme.of(context).colorScheme.secondary,
              ),
            ),
          Padding(
            padding: const EdgeInsets.only(left: 4.0),
            child: Text(
              room.lastMessage?.createdAt.localizedTimeShort(context) ?? room.createdAt.localizedTimeShort(context),
              style: TextStyle(
                fontSize: 13,
                color: unread
                    ? Theme.of(context).colorScheme.secondary
                    : Theme.of(context).textTheme.bodyText2!.color,
              ),
            ),
          ),
        ],
      ),
      subtitle: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          if (typingText.isEmpty &&
              ownMessage && false
              //room.lastMessage!.status.isSending
          ) ...[
            const SizedBox(
              width: 16,
              height: 16,
              child: CircularProgressIndicator.adaptive(strokeWidth: 2),
            ),
            const SizedBox(width: 4),
          ],
          AnimatedContainer(
            width: typingText.isEmpty ? 0 : 18,
            clipBehavior: Clip.hardEdge,
            decoration: const BoxDecoration(),
            duration: const Duration(milliseconds: 300),
            curve: Curves.bounceInOut,
            padding: const EdgeInsets.only(right: 4),
            child: Icon(
              Icons.edit_outlined,
              color: Theme.of(context).colorScheme.secondary,
              size: 14,
            ),
          ),
          Expanded(
            child: typingText.isNotEmpty
                ? Text(
                    typingText,
                    style: TextStyle(
                      color: Theme.of(context).colorScheme.secondary,
                    ),
                    softWrap: false,
                  )
                : Text(
                    //room.membership == Membership.invite
                    false
                        ? L10n.of(context)!.youAreInvitedToThisChat
                        : room.lastMessage == null ?
                            L10n.of(context)!.emptyChat : '${room.lastMessage!.senderId == StoreHelper.shared.userID
                        ? L10n.of(context)!.you
                        : (room.lastMessage!.sender.getFullName())}: ${room.lastMessage!.text}',
                    softWrap: false,
                    maxLines: 1,
                    overflow: TextOverflow.ellipsis,
                    style: TextStyle(
                      color: unread
                          ? Theme.of(context).colorScheme.secondary
                          : Theme.of(context).textTheme.bodyText2!.color,
                      decoration: false //room.lastMessage?.redacted == true
                          ? TextDecoration.lineThrough
                          : null,
                    ),
                  ),
          ),
          const SizedBox(width: 8),
          AnimatedContainer(
            duration: const Duration(milliseconds: 300),
            curve: Curves.bounceInOut,
            padding: const EdgeInsets.symmetric(horizontal: 7),
            height: unreadBubbleSize,
            width:
                room.unreadCount == 0 //&& !unread && false //!room.hasNewMessages
                    ? 0
                    : (unreadBubbleSize - 9) *
                            room.unreadCount.toString().length +
                        9,
            decoration: BoxDecoration(
              color: false //room.highlightCount > 0
                  ? Colors.red
                  : room.unreadCount > 0
                      ? Theme.of(context).primaryColor
                      : Theme.of(context).primaryColor.withAlpha(100),
              borderRadius: BorderRadius.circular(AppConfig.borderRadius),
            ),
            child: Center(
              child: room.unreadCount > 0
                  ? Text(
                      room.unreadCount.toString(),
                      style: const TextStyle(
                        color: Colors.white,
                        fontSize: 13,
                      ),
                    )
                  : Container(),
            ),
          ),
        ],
      ),
      onTap: () => clickAction(context),
    );
  }
}
