import os
from pathlib import Path
import polars
fpath = Path('./refgen.txt')
def format_keylist(lk):
    for i, v in enumerate(lk):
        listkeys[i] = f"{i}, {v}\n"
    return listkeys
def quickwrite(file,list):
    with file.open("w") as fo:
        fo.writelines(list)
        fo.close()

# listkeys builds output text for go
# with enumerated int that refers to key and actual key in comment
# Just enumerate this to get int that corresponds to key
listkeys = ["//KeyA","//KeyB","//KeyC","//KeyD","//KeyE","//KeyF",
            "//KeyG","//KeyH","//KeyI","//KeyJ","//KeyK","//KeyL",
            "//KeyM","//KeyN","//KeyO","//KeyP","//KeyQ","//KeyR",
            "//KeyS","//KeyT","//KeyU","//KeyV","//KeyW","//KeyX",
            "//KeyY","//KeyZ", # 0 to 25
            "//KeyAltLeft","//KeyAltRight",
            "//KeyArrowDown", "//KeyArrowLeft", "//KeyArrowRight", "//KeyArrowUp",
            "//KeyBackquote", "//KeyBackslash", "//KeyBackspace",
            "//KeyBracketLeft", "//KeyBracketRight",
            "//KeyCapsLock", # 37
            "//KeyComma",
            "//KeyContextMenu",
            "//KeyControlLeft", "//KeyControlRight",
            "//KeyDelete", # 42
            "//KeyDigit0", "//KeyDigit1", "//KeyDigit2", "//KeyDigit3", "//KeyDigit4",
            "//KeyDigit5", "//KeyDigit6", "//KeyDigit7", "//KeyDigit8", "//KeyDigit9",
            "//KeyEnd", "//KeyEnter", "//KeyEqual", "//KeyEscape", # 57

            "//KeyF1", "//KeyF2", "//KeyF3", "//KeyF4", "//KeyF5", "//KeyF6",
            "//KeyF7", "//KeyF8", "//KeyF9", "//KeyF10", "//KeyF11", "//KeyF12",
            "//KeyF13", "//KeyF14", "//KeyF15", "//KeyF16", "//KeyF17", "//KeyF18",
            "//KeyF19", "//KeyF20", "//KeyF21", "//KeyF22", "//KeyF23", "//KeyF24",

            "//KeyHome", "//KeyInsert",
            "//KeyIntlBackslas", "//KeyMetaLeft", "//KeyMetaRight",
            "//KeyMinus",
            "//KeyNumLock",
            "//KeyNumpad0", "//KeyNumpad1", "//KeyNumpad2", "//KeyNumpad3", "//KeyNumpad4",
            "//KeyNumpad5", "//KeyNumpad6", "//KeyNumpad7", "//KeyNumpad8", "//KeyNumpad9",
            "//KeyNumpadAdd", "//KeyNumpadDecima", "//KeyNumpadDivide", "//KeyNumpadEnter",
            "//KeyNumpadEqual", "//KeyNumpadMultip", "//KeyNumpadSubtra",
                "//KeyPageDown", "//KeyPageUp", "//KeyPause",
            "//KeyPeriod",
                "//KeyPrintScreen",
            "//KeyQuote",
                "//KeyScrollLock",
            "//KeySemicolon",
            "//KeyShiftLeft", "//KeyShiftRight",
            "//KeySlash", "//KeySpace", "//KeyTab",
            "//KeyAlt", "//KeyControl", "//KeyShift", "//KeyMeta", "//KeyMax"
    ]