import 'package:flutter/material.dart';

import 'package:flutter_gen/gen_l10n/l10n.dart';
import 'package:matrix/matrix.dart';
import 'package:matrix_link_text/link_text.dart';
import 'package:vrouter/vrouter.dart';
import 'package:cardio_messenger/utils/int_uri.dart';
import 'package:cardio_messenger/config/app_config.dart';
import 'package:cardio_messenger/pages/chat_details/chat_details.dart';
import 'package:cardio_messenger/pages/chat_details/participant_list_item.dart';
import 'package:cardio_messenger/utils/fluffy_share.dart';
import 'package:cardio_messenger/utils/matrix_sdk_extensions.dart/matrix_locals.dart';
import 'package:cardio_messenger/widgets/avatar.dart';
import 'package:cardio_messenger/widgets/chat_settings_popup_menu.dart';
import 'package:cardio_messenger/widgets/content_banner.dart';
import 'package:cardio_messenger/widgets/layouts/max_width_body.dart';
import 'package:cardio_messenger/widgets/matrix.dart';
import '../../utils/url_launcher.dart';

class ChatDetailsView extends StatelessWidget {
  final ChatDetailsController controller;

  const ChatDetailsView(this.controller, {Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final room = controller.room;
    if (room == null) {
      return Scaffold(
        appBar: AppBar(
          title: Text(L10n.of(context)!.loadingPleaseWait),
        ),
        body: Center(
          child: Text(L10n.of(context)!.loadingPleaseWait),
        ),
      );
    }

    //controller.members!.removeWhere((u) => u.membership == Membership.leave);
    // final actualMembersCount = (room.summary.mInvitedMemberCount ?? 0) +
    //     (room.summary.mJoinedMemberCount ?? 0);
    final actualMembersCount = 1;

    final canRequestMoreMembers =
        controller.members!.length < actualMembersCount;
    final iconColor = Theme.of(context).textTheme.bodyText1!.color;
    return StreamBuilder(
        stream: null,
        builder: (context, snapshot) {
          return Scaffold(
            body: NestedScrollView(
              headerSliverBuilder:
                  (BuildContext context, bool innerBoxIsScrolled) => <Widget>[
                SliverAppBar(
                  leading: IconButton(
                    icon: const Icon(Icons.close_outlined),
                    onPressed: () =>
                        VRouter.of(context).path.startsWith('/spaces/')
                            ? VRouter.of(context).pop()
                            : VRouter.of(context)
                                .toSegments(['rooms', controller.roomId!]),
                  ),
                  elevation: Theme.of(context).appBarTheme.elevation,
                  expandedHeight: 300.0,
                  floating: true,
                  pinned: true,
                  actions: <Widget>[
                    // if (room.name.isNotEmpty)
                    //   IconButton(
                    //     tooltip: L10n.of(context)!.share,
                    //     icon: Icon(Icons.adaptive.share_outlined),
                    //     onPressed: () => FluffyShare.share(
                    //         AppConfig.inviteLinkPrefix + room.name,
                    //         context),
                    //   ),
                    ChatSettingsPopupMenu(room, false)
                  ],
                  title: Text(
                      room.name,
                      style: TextStyle(
                          color: Theme.of(context)
                              .appBarTheme
                              .titleTextStyle!
                              .color)),
                  backgroundColor:
                      Theme.of(context).appBarTheme.backgroundColor,
                  flexibleSpace: FlexibleSpaceBar(
                    background: ContentBanner(
                        mxContent: room.avatarID?.uri,
                        onEdit: true //room.canSendEvent('m.room.avatar')
                            ? controller.setAvatarAction
                            : null),
                  ),
                ),
              ],
              body: MaxWidthBody(
                child: ListView.builder(
                  itemCount: controller.members!.length +
                      1 +
                      (canRequestMoreMembers ? 1 : 0),
                  itemBuilder: (BuildContext context, int i) => i == 0
                      ? Column(
                          crossAxisAlignment: CrossAxisAlignment.stretch,
                          children: <Widget>[
                            // ListTile(
                            //   leading: true//room.canSendEvent('m.room.topic')
                            //       ? CircleAvatar(
                            //           backgroundColor: Theme.of(context)
                            //               .scaffoldBackgroundColor,
                            //           foregroundColor: iconColor,
                            //           radius: Avatar.defaultSize / 2,
                            //           child: const Icon(Icons.edit_outlined),
                            //         )
                            //       : null,
                            //   title: Text(
                            //       '${L10n.of(context)!.groupDescription}:',
                            //       style: TextStyle(
                            //           color: Theme.of(context)
                            //               .colorScheme
                            //               .secondary,
                            //           fontWeight: FontWeight.bold)),
                            //   subtitle: LinkText(
                            //     text: room.about.isEmpty
                            //         ? L10n.of(context)!.addGroupDescription
                            //         : room.about,
                            //     linkStyle:
                            //         const TextStyle(color: Colors.blueAccent),
                            //     textStyle: TextStyle(
                            //       fontSize: 14,
                            //       color: Theme.of(context)
                            //           .textTheme
                            //           .bodyText2!
                            //           .color,
                            //     ),
                            //     onLinkTap: (url) =>
                            //         UrlLauncher(context, url).launchUrl(),
                            //   ),
                            //   onTap: true //room.canSendEvent('m.room.topic')
                            //       ? controller.setTopicAction
                            //       : null,
                            // ),
                            //const SizedBox(height: 8),
                            //const Divider(height: 1),
                            // ListTile(
                            //   title: Text(
                            //     L10n.of(context)!.settings,
                            //     style: TextStyle(
                            //       color:
                            //           Theme.of(context).colorScheme.secondary,
                            //       fontWeight: FontWeight.bold,
                            //     ),
                            //   ),
                            //   trailing: Icon(controller.displaySettings
                            //       ? Icons.keyboard_arrow_down_outlined
                            //       : Icons.keyboard_arrow_right_outlined),
                            //   onTap: controller.toggleDisplaySettings,
                            // ),
                            if (controller.displaySettings) ...[
                              if (true)//room.canSendEvent('m.room.name'))
                                ListTile(
                                  leading: CircleAvatar(
                                    backgroundColor: Theme.of(context)
                                        .scaffoldBackgroundColor,
                                    foregroundColor: iconColor,
                                    child: const Icon(
                                        Icons.people_outline_outlined),
                                  ),
                                  title: Text(L10n.of(context)!
                                      .changeTheNameOfTheGroup),
                                  subtitle: Text(room.name),
                                  onTap: controller.setDisplaynameAction,
                                ),
                              // if (room.joinRules == JoinRules.public))
                              //   ListTile(
                              //     leading: CircleAvatar(
                              //       backgroundColor: Theme.of(context)
                              //           .scaffoldBackgroundColor,
                              //       foregroundColor: iconColor,
                              //       child: const Icon(Icons.link_outlined),
                              //     ),
                              //     onTap: controller.editAliases,
                              //     title:
                              //         Text(L10n.of(context)!.editRoomAliases),
                              //     subtitle: Text(
                              //         (room.name.isNotEmpty)
                              //             ? room.name
                              //             : L10n.of(context)!.none),
                              //   ),
                              // code200
                              // ListTile(
                              //   leading: CircleAvatar(
                              //     backgroundColor:
                              //         Theme.of(context).scaffoldBackgroundColor,
                              //     foregroundColor: iconColor,
                              //     child: const Icon(
                              //         Icons.insert_emoticon_outlined),
                              //   ),
                              //   title: Text(L10n.of(context)!.emoteSettings),
                              //   subtitle:
                              //       Text(L10n.of(context)!.setCustomEmotes),
                              //   onTap: controller.goToEmoteSettings,
                              // ),
                              // PopupMenuButton(
                              //   onSelected: controller.setJoinRulesAction,
                              //   itemBuilder: (BuildContext context) =>
                              //       <PopupMenuEntry<JoinRules>>[
                              //     if (room.canChangeJoinRules)
                              //       PopupMenuItem<JoinRules>(
                              //         value: JoinRules.public,
                              //         child: Text(JoinRules.public
                              //             .getLocalizedString(
                              //                 MatrixLocals(L10n.of(context)!))),
                              //       ),
                              //     if (room.canChangeJoinRules)
                              //       PopupMenuItem<JoinRules>(
                              //         value: JoinRules.invite,
                              //         child: Text(JoinRules.invite
                              //             .getLocalizedString(
                              //                 MatrixLocals(L10n.of(context)!))),
                              //       ),
                              //   ],
                              //   child: ListTile(
                              //     leading: CircleAvatar(
                              //         backgroundColor: Theme.of(context)
                              //             .scaffoldBackgroundColor,
                              //         foregroundColor: iconColor,
                              //         child: const Icon(Icons.shield_outlined)),
                              //     title: Text(L10n.of(context)!
                              //         .whoIsAllowedToJoinThisGroup),
                              //     subtitle: Text(
                              //       room.joinRules!.getLocalizedString(
                              //           MatrixLocals(L10n.of(context)!)),
                              //     ),
                              //   ),
                              // ),
                              // PopupMenuButton(
                              //   onSelected:
                              //       controller.setHistoryVisibilityAction,
                              //   itemBuilder: (BuildContext context) =>
                              //       <PopupMenuEntry<HistoryVisibility>>[
                              //     if (room.canChangeHistoryVisibility)
                              //       PopupMenuItem<HistoryVisibility>(
                              //         value: HistoryVisibility.invited,
                              //         child: Text(HistoryVisibility.invited
                              //             .getLocalizedString(
                              //                 MatrixLocals(L10n.of(context)!))),
                              //       ),
                              //     if (room.canChangeHistoryVisibility)
                              //       PopupMenuItem<HistoryVisibility>(
                              //         value: HistoryVisibility.joined,
                              //         child: Text(HistoryVisibility.joined
                              //             .getLocalizedString(
                              //                 MatrixLocals(L10n.of(context)!))),
                              //       ),
                              //     if (room.canChangeHistoryVisibility)
                              //       PopupMenuItem<HistoryVisibility>(
                              //         value: HistoryVisibility.shared,
                              //         child: Text(HistoryVisibility.shared
                              //             .getLocalizedString(
                              //                 MatrixLocals(L10n.of(context)!))),
                              //       ),
                              //     if (room.canChangeHistoryVisibility)
                              //       PopupMenuItem<HistoryVisibility>(
                              //         value: HistoryVisibility.worldReadable,
                              //         child: Text(HistoryVisibility
                              //             .worldReadable
                              //             .getLocalizedString(
                              //                 MatrixLocals(L10n.of(context)!))),
                              //       ),
                              //   ],
                              //   child: ListTile(
                              //     leading: CircleAvatar(
                              //       backgroundColor: Theme.of(context)
                              //           .scaffoldBackgroundColor,
                              //       foregroundColor: iconColor,
                              //       child:
                              //           const Icon(Icons.visibility_outlined),
                              //     ),
                              //     title: Text(L10n.of(context)!
                              //         .visibilityOfTheChatHistory),
                              //     subtitle: Text(
                              //       room.historyVisibility!.getLocalizedString(
                              //           MatrixLocals(L10n.of(context)!)),
                              //     ),
                              //   ),
                              // ),
                              // if (room.joinRules == JoinRules.public)
                              //   PopupMenuButton(
                              //     onSelected: controller.setGuestAccessAction,
                              //     itemBuilder: (BuildContext context) =>
                              //         <PopupMenuEntry<GuestAccess>>[
                              //       if (room.canChangeGuestAccess)
                              //         PopupMenuItem<GuestAccess>(
                              //           value: GuestAccess.canJoin,
                              //           child: Text(
                              //             GuestAccess.canJoin
                              //                 .getLocalizedString(MatrixLocals(
                              //                     L10n.of(context)!)),
                              //           ),
                              //         ),
                              //       if (room.canChangeGuestAccess)
                              //         PopupMenuItem<GuestAccess>(
                              //           value: GuestAccess.forbidden,
                              //           child: Text(
                              //             GuestAccess.forbidden
                              //                 .getLocalizedString(MatrixLocals(
                              //                     L10n.of(context)!)),
                              //           ),
                              //         ),
                              //     ],
                              //     child: ListTile(
                              //       leading: CircleAvatar(
                              //         backgroundColor: Theme.of(context)
                              //             .scaffoldBackgroundColor,
                              //         foregroundColor: iconColor,
                              //         child: const Icon(
                              //             Icons.person_add_alt_1_outlined),
                              //       ),
                              //       title: Text(L10n.of(context)!
                              //           .areGuestsAllowedToJoin),
                              //       subtitle: Text(
                              //         room.guestAccess.getLocalizedString(
                              //             MatrixLocals(L10n.of(context)!)),
                              //       ),
                              //     ),
                              //   ),
                              // ListTile(
                              //   title:
                              //       Text(L10n.of(context)!.editChatPermissions),
                              //   subtitle: Text(
                              //       L10n.of(context)!.whoCanPerformWhichAction),
                              //   leading: CircleAvatar(
                              //     backgroundColor:
                              //         Theme.of(context).scaffoldBackgroundColor,
                              //     foregroundColor: iconColor,
                              //     child: const Icon(
                              //         Icons.edit_attributes_outlined),
                              //   ),
                              //   onTap: () =>
                              //       VRouter.of(context).to('permissions'),
                              // ),
                            ],
                            // const Divider(height: 1),
                            // ListTile(
                            //   title: Text(
                            //     actualMembersCount > 1
                            //         ? L10n.of(context)!.countParticipants(
                            //             actualMembersCount.toString())
                            //         : L10n.of(context)!.emptyChat,
                            //     style: TextStyle(
                            //       color:
                            //           Theme.of(context).colorScheme.secondary,
                            //       fontWeight: FontWeight.bold,
                            //     ),
                            //   ),
                            // ),
                            //room.canInvite
                            true
                                ? ListTile(
                                    title:
                                        Text(L10n.of(context)!.inviteContact),
                                    leading: CircleAvatar(
                                      backgroundColor:
                                          Theme.of(context).primaryColor,
                                      foregroundColor: Colors.white,
                                      radius: Avatar.defaultSize / 2,
                                      child: const Icon(Icons.add_outlined),
                                    ),
                                    onTap: () =>
                                        VRouter.of(context).to('invite'),
                                  )
                                : Container(),
                          ],
                        )
                      : i < controller.members!.length + 1
                          ? ParticipantListItem(controller.members![i - 1], room)
                          : ListTile(
                              title: Text(L10n.of(context)!
                                  .loadCountMoreParticipants(
                                      (actualMembersCount -
                                              controller.members!.length)
                                          .toString())),
                              leading: CircleAvatar(
                                backgroundColor:
                                    Theme.of(context).scaffoldBackgroundColor,
                                child: const Icon(
                                  Icons.refresh,
                                  color: Colors.grey,
                                ),
                              ),
                              onTap: controller.requestMoreMembersAction,
                            ),
                ),
              ),
            ),
          );
        });
  }
}
