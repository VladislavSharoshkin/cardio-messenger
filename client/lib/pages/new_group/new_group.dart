import 'package:cardio_messenger/services/get_chat.dart';
import 'package:flutter/material.dart';

import 'package:future_loading_dialog/future_loading_dialog.dart';
import 'package:matrix/matrix.dart' as sdk;
import 'package:vrouter/vrouter.dart';

import 'package:cardio_messenger/pages/new_group/new_group_view.dart';
import 'package:cardio_messenger/widgets/matrix.dart';

class NewGroup extends StatefulWidget {
  const NewGroup({Key? key}) : super(key: key);

  @override
  NewGroupController createState() => NewGroupController();
}

class NewGroupController extends State<NewGroup> {
  TextEditingController controller = TextEditingController();
  bool publicGroup = false;

  void setPublicGroup(bool b) => setState(() => publicGroup = b);

  void submitAction([_]) async {

    final roomID = await showFutureLoadingDialog(
      context: context,
      future: () async {
        var getChat = await GetChat.createGroupChat(controller.text.isNotEmpty ? controller.text : null, []);
        int roomId = getChat.chat.id;
        return roomId;
      },
    );
    if (roomID.error == null) {
      VRouter.of(context).toSegments(['rooms', roomID.result!.toString(), 'invite']);
    }
  }

  @override
  Widget build(BuildContext context) => NewGroupView(this);
}
