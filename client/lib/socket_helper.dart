import 'package:cardio_messenger/store_helper.dart';
import 'package:socket_io_client/socket_io_client.dart';

class SocketHelper {
  static final shared = SocketHelper();
  final Socket socket = io('https://${StoreHelper.shared.domain?.split(':').first}:27992', <String, dynamic>{
    'transports': ['websocket'],
    'autoConnect': false,
    'forceNew':true
  });

  void disconnectSocket(){
    socket.disconnect();
  }
  void connectSocket() async {
    if (socket.connected) {
      return;
    }
    print("connect to socket");
    socket.connect();
    socket.onDisconnect((data) => (data) {
      print(data);
    });
    socket.on('connect', (data) async {
      print('socket connected');
      socket.emit('token', StoreHelper.shared.token);

      socket.on('update', (data) {
        print('socket update');
      });

    });
  }
}
