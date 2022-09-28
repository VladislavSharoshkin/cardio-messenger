import 'package:cardio_messenger/models/filee.dart';
import 'package:cardio_messenger/services/get_file.dart';
import 'package:flutter/material.dart';

import 'package:matrix/matrix.dart';
import 'package:open_file/open_file.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:vrouter/vrouter.dart';

import 'package:cardio_messenger/pages/image_viewer/image_viewer_view.dart';
import 'package:cardio_messenger/utils/platform_infos.dart';
import 'package:cardio_messenger/widgets/matrix.dart';
import '../../utils/matrix_sdk_extensions.dart/event_extension.dart';

class ImageViewer extends StatefulWidget {
  final MFile event;
  final void Function()? onLoaded;

  const ImageViewer(this.event, {Key? key, this.onLoaded}) : super(key: key);

  @override
  ImageViewerController createState() => ImageViewerController();
}

class ImageViewerController extends State<ImageViewer> {
  /// Forward this image to another room.
  void forwardAction() {
    //Matrix.of(context).shareContent = widget.event.content;
    VRouter.of(context).to('/rooms');
  }

  /// Save this file with a system call.
  void saveFileAction() => GetFile.open(widget.event);

  static const maxScaleFactor = 1.5;

  /// Go back if user swiped it away
  void onInteractionEnds(ScaleEndDetails endDetails) {
    if (PlatformInfos.usesTouchscreen == false) {
      if (endDetails.velocity.pixelsPerSecond.dy >
          MediaQuery.of(context).size.height * maxScaleFactor) {
        Navigator.of(context, rootNavigator: false).pop();
      }
    }
  }

  @override
  Widget build(BuildContext context) => ImageViewerView(this);
}
