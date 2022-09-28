import 'package:cardio_messenger/models/chat.dart';
import 'package:cardio_messenger/models/message.dart';
import 'package:cardio_messenger/pages/chat/events/attachment.dart';
import 'package:cardio_messenger/store_helper.dart';
import 'package:flutter/material.dart';

import 'package:matrix/matrix.dart';
import 'package:cardio_messenger/utils/int_uri.dart';
import 'package:cardio_messenger/config/themes.dart';
import 'package:cardio_messenger/utils/date_time_extension.dart';
import 'package:cardio_messenger/utils/string_color.dart';
import 'package:cardio_messenger/widgets/avatar.dart';
import 'package:cardio_messenger/widgets/matrix.dart';
import '../../../config/app_config.dart';
import 'message_content.dart';
import 'message_reactions.dart';
import 'reply_content.dart';
import 'state_message.dart';
import 'verification_request_content.dart';

class Message extends StatelessWidget {
  final MMessage event;
  final MMessage? nextEvent;
  final void Function(MMessage)? onSelect;
  final void Function(Event)? onAvatarTab;
  final void Function(Event)? onInfoTab;
  final void Function(String)? scrollToEventId;
  final void Function(int) unfold;
  final bool longPressSelect;
  final bool selected;
  final Timeline? timeline;

  const Message(this.event,
      {this.nextEvent,
      this.longPressSelect = false,
      this.onSelect,
      this.onInfoTab,
      this.onAvatarTab,
      this.scrollToEventId,
      required this.unfold,
      this.selected = false,
      this.timeline,
      Key? key})
      : super(key: key);

  /// Indicates wheither the user may use a mouse instead
  /// of touchscreen.
  static bool useMouse = false;

  @override
  Widget build(BuildContext context) {
    // if (!{
    //   EventTypes.Message,
    //   EventTypes.Sticker,
    //   EventTypes.Encrypted,
    //   EventTypes.CallInvite
    // }.contains(event.type)) {
    //   if (event.type.startsWith('m.call.')) {
    //     return Container();
    //   }
    //   return StateMessage(event, unfold: unfold);
    // }
    //
    // if (event.type == EventTypes.Message &&
    //     event.messageType == EventTypes.KeyVerificationRequest) {
    //   return VerificationRequestContent(event: event, timeline: timeline);
    // }

    final client = Matrix.of(context).client;
    final ownMessage = event.senderId == StoreHelper.shared.userID;
    final alignment = ownMessage ? Alignment.topRight : Alignment.topLeft;
    var color = Theme.of(context).appBarTheme.backgroundColor;
    final displayTime = //event.type == EventTypes.RoomCreate ||
        nextEvent == null ||
            !event.createdAt.sameEnvironment(nextEvent!.createdAt);

    final sameSender = nextEvent != null &&
            [
              EventTypes.Message,
              EventTypes.Sticker,
              EventTypes.Encrypted,
            ].contains(nextEvent!.type)
        ? nextEvent!.sender.id == event.sender.id && !displayTime
        : false;
    final textColor = ownMessage
        ? Colors.white
        : Theme.of(context).brightness == Brightness.dark
            ? Colors.white
            : Colors.black;
    final rowMainAxisAlignment =
        ownMessage ? MainAxisAlignment.end : MainAxisAlignment.start;

    MMessage displayEvent = event; //event.getDisplayEvent(timeline);
    final borderRadius = BorderRadius.only(
      topLeft: !ownMessage
          ? const Radius.circular(2)
          : const Radius.circular(AppConfig.borderRadius),
      topRight: ownMessage
          ? const Radius.circular(2)
          : const Radius.circular(AppConfig.borderRadius),
      bottomLeft: const Radius.circular(AppConfig.borderRadius),
      bottomRight: const Radius.circular(AppConfig.borderRadius),
    );
    final noBubble = {
      MessageTypes.Video,
      MessageTypes.Image,
      MessageTypes.Sticker,
    }.contains(event.type);

    if (ownMessage) {
      // color = displayEvent.status.isError
      //     ? Colors.redAccent
      //     : Theme.of(context).primaryColor;
      color = Theme.of(context).primaryColor;
    }

    final rowChildren = <Widget>[
      sameSender || ownMessage
          ? SizedBox(
              width: Avatar.defaultSize,
              child: Padding(
                padding: const EdgeInsets.only(top: 8.0),
                child: Center(
                  child: SizedBox(
                    width: 16 * AppConfig.bubbleSizeFactor,
                    height: 16 * AppConfig.bubbleSizeFactor,
                    // child: event.status == EventStatus.sending
                    //     ? const CircularProgressIndicator.adaptive(
                    //         strokeWidth: 2)
                    //     : event.status == EventStatus.error
                    //         ? const Icon(Icons.error, color: Colors.red)
                    //         : null,
                  ),
                ),
              ))
          : Avatar(
              mxContent: event.sender.avatarID?.uri,
              name: event.sender.getFullName(),
              onTap: () => 'onAvatarTab!(event)',
            ),
      Expanded(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            if (!sameSender)
              Padding(
                padding: const EdgeInsets.only(left: 8.0, bottom: 4),
                child: ownMessage // || event.room.isDirectChat
                    ? const SizedBox(height: 12)
                    : Text(
                        event.sender.getFullName(),
                        style: TextStyle(
                          fontSize: 12,
                          fontWeight: FontWeight.bold,
                          color: event.sender.getFullName().color,
                        ),
                      ),
              ),
            Container(
              alignment: alignment,
              padding: const EdgeInsets.only(left: 8),
              child: Material(
                color: noBubble ? Colors.transparent : color,
                elevation: event.type == EventTypes.Sticker ? 0 : 6,
                shadowColor:
                    Theme.of(context).secondaryHeaderColor.withAlpha(100),
                borderRadius: borderRadius,
                clipBehavior: Clip.antiAlias,
                child: InkWell(
                  onHover: (b) => useMouse = true,
                  onTap: !useMouse && longPressSelect
                      ? () {}
                      : () => onSelect!(event),
                  onLongPress: !longPressSelect ? null : () => onSelect!(event),
                  borderRadius: borderRadius,
                  child: Container(
                    decoration: BoxDecoration(
                      borderRadius:
                          BorderRadius.circular(AppConfig.borderRadius),
                    ),
                    padding: noBubble
                        ? EdgeInsets.zero
                        : EdgeInsets.all(16 * AppConfig.bubbleSizeFactor),
                    constraints: const BoxConstraints(
                        maxWidth: FluffyThemes.columnWidth * 1.5),
                    child: Stack(
                      children: <Widget>[
                        Column(
                          mainAxisSize: MainAxisSize.min,
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: <Widget>[
                            // if (event.relationshipType ==
                            //     RelationshipTypes.reply)
                            //   FutureBuilder<Event?>(
                            //     future: event.getReplyEvent(timeline),
                            //     builder: (BuildContext context, snapshot) {
                            //       final replyEvent = snapshot.hasData
                            //           ? snapshot.data!
                            //           : Event(
                            //               eventId: event.relationshipEventId!,
                            //               content: {
                            //                 'msgtype': 'm.text',
                            //                 'body': '...'
                            //               },
                            //               senderId: event.senderId,
                            //               type: 'm.room.message',
                            //               room: event.room,
                            //               status: EventStatus.sent,
                            //               originServerTs: DateTime.now(),
                            //             );
                            //       return InkWell(
                            //         onTap: () {
                            //           if (scrollToEventId != null) {
                            //             scrollToEventId!(replyEvent.eventId);
                            //           }
                            //         },
                            //         child: AbsorbPointer(
                            //           child: Container(
                            //             margin: EdgeInsets.symmetric(
                            //                 vertical: 4.0 *
                            //                     AppConfig.bubbleSizeFactor),
                            //             child: ReplyContent(replyEvent,
                            //                 lightText: ownMessage,
                            //                 timeline: timeline),
                            //           ),
                            //         ),
                            //       );
                            //     },
                            //   ),
                            MessageContent(
                              displayEvent,
                              textColor: textColor,
                              onInfoTab: onInfoTab,
                            ),

                            if (displayEvent.forwardMessages.isNotEmpty)
                              Padding(
                                padding: EdgeInsets.only(
                                    top: 4.0 * AppConfig.bubbleSizeFactor),
                                child: Row(
                                  mainAxisSize: MainAxisSize.min,
                                  children: [
                                    Icon(
                                      Icons.adaptive.share,
                                      color: textColor.withAlpha(164),
                                      size: 14,
                                    ),
                                    Text(
                                      ' переслано:',
                                      style: TextStyle(
                                        color: textColor.withAlpha(164),
                                        fontSize: 12,
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                            for (var forwardMessage
                                in displayEvent.forwardMessages)
                              MessageContent(
                                forwardMessage.forwardMessage,
                                textColor: textColor,
                              ),
                            // if (event.hasAggregatedEvents(
                            //     timeline, RelationshipTypes.edit))
                            //   Padding(
                            //     padding: EdgeInsets.only(
                            //         top: 4.0 * AppConfig.bubbleSizeFactor),
                            //     child: Row(
                            //       mainAxisSize: MainAxisSize.min,
                            //       children: [
                            //         Icon(
                            //           Icons.edit_outlined,
                            //           color: textColor.withAlpha(164),
                            //           size: 14,
                            //         ),
                            //         Text(
                            //           ' - ${displayEvent.originServerTs.localizedTimeShort(context)}',
                            //           style: TextStyle(
                            //             color: textColor.withAlpha(164),
                            //             fontSize: 12,
                            //           ),
                            //         ),
                            //       ],
                            //     ),
                            //   ),
                            if (event.sender.id == StoreHelper.shared.userID && event.readState == 1)
                              Padding(
                                padding: EdgeInsets.only(
                                    top: 4.0 * AppConfig.bubbleSizeFactor),
                                child: Row(
                                  mainAxisSize: MainAxisSize.min,
                                  children: [
                                    Icon(
                                      Icons.check,
                                      color: textColor.withAlpha(164),
                                      size: 14,
                                    ),
                                    Text(
                                      ' - доставлено',
                                      style: TextStyle(
                                        color: textColor.withAlpha(164),
                                        fontSize: 12,
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                          ],
                        ),
                      ],
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    ];
    final row = Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      mainAxisAlignment: rowMainAxisAlignment,
      children: rowChildren,
    );
    Widget container = row;
    if ( //event.hasAggregatedEvents(timeline, RelationshipTypes.reaction) ||
        displayTime || selected) {
      container = Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment:
            ownMessage ? CrossAxisAlignment.end : CrossAxisAlignment.start,
        children: <Widget>[
          if (displayTime || selected)
            Padding(
              padding: displayTime
                  ? EdgeInsets.symmetric(
                      vertical: 8.0 * AppConfig.bubbleSizeFactor)
                  : EdgeInsets.zero,
              child: Center(
                  child: Material(
                color: displayTime
                    ? Theme.of(context).backgroundColor
                    : Theme.of(context).backgroundColor.withOpacity(0.33),
                borderRadius: BorderRadius.circular(AppConfig.borderRadius / 2),
                clipBehavior: Clip.antiAlias,
                child: Padding(
                  padding: const EdgeInsets.all(6.0),
                  child: Text(
                    event.createdAt.localizedTime(context),
                    style: TextStyle(fontSize: 14 * AppConfig.fontSizeFactor),
                  ),
                ),
              )),
            ),
          row,
          // if (event.hasAggregatedEvents(timeline, RelationshipTypes.reaction))
          //   Padding(
          //     padding: EdgeInsets.only(
          //       top: 4.0 * AppConfig.bubbleSizeFactor,
          //       left: (ownMessage ? 0 : Avatar.defaultSize) + 12.0,
          //       right: 12.0,
          //     ),
          //     child: MessageReactions(event, timeline),
          //   ),
        ],
      );
    } else {
      container = row;
    }

    return Center(
      child: Container(
        color: selected
            ? Theme.of(context).primaryColor.withAlpha(100)
            : Theme.of(context).primaryColor.withAlpha(0),
        constraints:
            const BoxConstraints(maxWidth: FluffyThemes.columnWidth * 2.5),
        child: Padding(
          padding: EdgeInsets.symmetric(
            horizontal: 8.0,
            vertical: 4.0 * AppConfig.bubbleSizeFactor,
          ),
          child: container,
        ),
      ),
    );
  }
}
