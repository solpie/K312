#ifndef MAINC
#define MAINC

#include <stdio.h>
#include <windows.h>

///// C call go

extern int onKey(void *, void *);
extern void log1(void *);   

////////////////  go call C 
void keyEvent(int state,int keyCode){
    keybd_event(keyCode, 0, state, 0);
}
#define MOUSEEVENTF_MOVE        0x00000001
#define MOUSEEVENTF_LEFTDOWN    0x00000002
#define MOUSEEVENTF_LEFTUP      0x00000004
#define MOUSEEVENTF_RIGHTDOWN   0x00000008
#define MOUSEEVENTF_RIGHTUP     0x00000010
#define MOUSEEVENTF_MIDDLEDOWN  0x00000020
#define MOUSEEVENTF_MIDDLEUP    0x00000040
void mouseEvent(int event){
    mouse_event(event, 0, 0, 0, 0);
}

char *getActTitle(){
    char title[256];
    GetWindowText(GetForegroundWindow(),title,256);
    // printf("%s\n",title);
    // log1(title);
    return title;
}
///////////////////////////////////

__declspec(dllexport)
LRESULT CALLBACK onKeyEvent(int code, WPARAM wparam, LPARAM lparam) {
    KBDLLHOOKSTRUCT keyInfo;
    if (code == HC_ACTION
        && (wparam == WM_SYSKEYUP || wparam == WM_KEYUP || wparam == WM_SYSKEYDOWN
            || wparam == WM_KEYDOWN)) {
        keyInfo = *((KBDLLHOOKSTRUCT *) lparam);

        if (keyInfo.flags & LLKHF_INJECTED) {    // process injected events like normal, because most probably we are injecting them
            //logKeyEvent("injected", keyInfo);
            return CallNextHookEx(NULL, code, wparam, lparam);
        }
    }
    if (code == HC_ACTION && (wparam == WM_SYSKEYUP || wparam == WM_KEYUP)) {
        if (onKey(2, keyInfo.vkCode))
            return -1;
    }
    else if (code == HC_ACTION && (wparam == WM_SYSKEYDOWN || wparam == WM_KEYDOWN)) {
        if (onKey(1, keyInfo.vkCode))
            return -1;
    }
    return CallNextHookEx(NULL, code, wparam, lparam);
}


DWORD WINAPI hookKeyThread(void *user) {
    HINSTANCE base = GetModuleHandle(NULL);
    MSG msg;
    printf("[hookKeyThread] start!!!\n");
    if (!base) {
        if (!(base = LoadLibrary((wchar_t *) user))) {
            return 1;
        }
    }

    HWND keyHook = SetWindowsHookEx(WH_KEYBOARD_LL, onKeyEvent, base, 0);

    while (GetMessage(&msg, 0, 0, 0) > 0) {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }

    UnhookWindowsHookEx(keyHook);
    return 0;
}

void hook() {
//    InitCommonControls();
    HANDLE thread = CreateThread(0, 0, hookKeyThread, NULL, 0, NULL);
    MSG msg;
    while (GetMessage(&msg, NULL, 0, 0) > 0) {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }
}

#endif

