This app aim to impelment asymmetry encryption via QRcode on messaging app like line/telegram.

Workflow:
    1. Implement asymmetric encryption to message
        a. Key-gen
        b. Encrypt message
        c. Decrypt message
    2. Implement messaging functions
        a. Send message (bind with encryption)
        b. recieve message
    3. Impement QRCode generation
        a. Translate message to QRCode
        b. Translate QRCode to message
    4. OS implementation
        a. Read QRCode from screen (send to files or buffer)
        b. Send QRCode to app
