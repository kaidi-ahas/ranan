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

        String noteStr = message.substring(0, commaIndex);
        float cents    = message.substring(commaIndex + 1).toFloat();

        digitalWrite(PIN_LED_LEFT,  LOW);
        digitalWrite(PIN_LED_GREEN, LOW);
        digitalWrite(PIN_LED_RIGHT, LOW);

        if (cents < -TUNE_THRESHOLD) {
            digitalWrite(PIN_LED_LEFT, HIGH);
        } else if (cents > TUNE_THRESHOLD) {
            digitalWrite(PIN_LED_RIGHT, HIGH);
        } else {
            digitalWrite(PIN_LED_GREEN, HIGH);
        }
    }
}