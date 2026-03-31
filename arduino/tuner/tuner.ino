const int PIN_LED_LEFT  = 10;
const int PIN_LED_GREEN = 11;
const int PIN_LED_RIGHT = 12;

const float TUNE_THRESHOLD = 10.0;

void setup() {
    Serial.begin(9600);
    pinMode(PIN_LED_LEFT,  OUTPUT);
    pinMode(PIN_LED_GREEN, OUTPUT);
    pinMode(PIN_LED_RIGHT, OUTPUT);
}

void loop() {
    if (Serial.available() > 0) {
        String message = Serial.readStringUntil('\n');
        message.trim();

        int commaIndex = message.indexOf(',');
        if (commaIndex == -1) {
            return;
        }

        String status = message.substring(commaIndex +1);

        digitalWrite(PIN_LED_LEFT,  LOW);
        digitalWrite(PIN_LED_GREEN, LOW);
        digitalWrite(PIN_LED_RIGHT, LOW);

        if (status == "flat") {
            digitalWrite(PIN_LED_LEFT, HIGH);
        } else if (status == "sharp") {
            digitalWrite(PIN_LED_RIGHT, HIGH);
        } else if (status == "intune") {
            digitalWrite(PIN_LED_GREEN, HIGH);
        }
    }
}