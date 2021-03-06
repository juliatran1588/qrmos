import "package:flutter/material.dart";
import 'package:qrmos/services/qrmos/qrmos.dart' as qrmos_api;

class LoginScreen extends StatefulWidget {
  const LoginScreen({Key? key}) : super(key: key);

  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  String _username = "";
  String _password = "";
  String _errMsg = "";

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Login"),
      ),
      body: Center(
        child: Column(mainAxisAlignment: MainAxisAlignment.center, children: <Widget>[
          TextFormField(
            decoration: _inputDecor("Username"),
            onChanged: _onUsernameChanged,
            keyboardType: TextInputType.text,
          ),
          TextFormField(
            decoration: _inputDecor("Password"),
            onChanged: _onPasswordChanges,
            obscureText: true,
          ),
          Text(_errMsg, style: TextStyle(color: Theme.of(context).errorColor)),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: ElevatedButton(
              onPressed: _onLoginButtonPressed(context),
              child: const Text("Login"),
            ),
          ),
        ]),
      ),
    );
  }

  InputDecoration _inputDecor(String labelText) {
    return InputDecoration(
      labelText: labelText,
      constraints: const BoxConstraints(maxWidth: 300),
    );
  }

  void _onUsernameChanged(String username) {
    _username = username;
    setState(() {
      _errMsg = "";
    });
  }

  void _onPasswordChanges(String password) {
    _password = password;
    setState(() {
      _errMsg = "";
    });
  }

  void Function() _onLoginButtonPressed(BuildContext context) {
    return () async {
      var loginResp = await qrmos_api.login(_username, _password);
      if (loginResp.error != null) {
        setState(() {
          _errMsg = "Error: " + loginResp.error!.message;
        });
        return;
      }
      Navigator.of(context).pop();
    };
  }
}
