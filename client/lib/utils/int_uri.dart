import 'package:cardio_messenger/store_helper.dart';
import 'package:flutter/material.dart';

extension IntUri on int {
  Uri get uri {
    return Uri.parse('https://${StoreHelper.shared.domain}/file/download/$this');
  }

  String get sizeString {

      num size = this;
      if (size < 1000000) {
        size = size / 1000;
        size = (size * 10).round() / 10;
        return '${size.toString()} KB';
      } else if (size < 1000000000) {
        size = size / 1000000;
        size = (size * 10).round() / 10;
        return '${size.toString()} MB';
      } else {
        size = size / 1000000000;
        size = (size * 10).round() / 10;
        return '${size.toString()} GB';
      }

  }
}