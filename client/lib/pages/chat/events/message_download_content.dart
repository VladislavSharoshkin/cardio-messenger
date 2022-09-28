import 'package:cardio_messenger/models/filee.dart';
import 'package:cardio_messenger/services/get_file.dart';
import 'package:cardio_messenger/utils/int_uri.dart';
import 'package:flutter/material.dart';

import 'package:matrix/matrix.dart';

import 'package:cardio_messenger/utils/matrix_sdk_extensions.dart/event_extension.dart';
import 'package:open_file/open_file.dart';
import 'package:url_launcher/url_launcher.dart';

class MessageDownloadContent extends StatelessWidget {
  final MFile event;
  final Color textColor;

  const MessageDownloadContent(this.event, this.textColor, {Key? key})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    final filename = event.name;//event.content.tryGet<String>('filename') ?? event.body;
    final filetype = (filename.contains('.')
        ? filename.split('.').last.toUpperCase()
        : 'UNKNOWN');
    final sizeString = event.size.sizeString;
    return InkWell(
      onTap: () => GetFile.open(event),//OpenFile.open(event.getURI().toString()),//event.saveFile(context),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          Row(
            children: [
              Icon(
                Icons.file_download_outlined,
                color: textColor,
              ),
              const SizedBox(width: 8),
              Expanded(
                child: Text(
                  filename,
                  maxLines: 1,
                  style: TextStyle(
                    color: textColor,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ],
          ),
          const Divider(),
          Row(
            children: [
              Text(
                filetype,
                style: TextStyle(
                  color: textColor.withAlpha(150),
                ),
              ),
              const Spacer(),
              if (sizeString != null)
                Text(
                  sizeString,
                  style: TextStyle(
                    color: textColor.withAlpha(150),
                  ),
                ),
            ],
          ),
        ],
      ),
    );
  }
}
