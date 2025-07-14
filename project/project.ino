#include "constants.h"

#include <Wire.h>
#include <LiquidCrystal_I2C.h>


/* values for display printing */
float last_update = 0.0f;
const float cool_down = 200.0f;    // update time of display

/* speed & deltatime */
float speed = 0.0f;
float acceleration = 0.0f;
float deltatime = 0.0f;
float last_time = 0.0f;

/* values of pins */ 
int pot_value   = 0;
int gas_state   = 0;
int break_state = 0;

int last_pot = 0;
LiquidCrystal_I2C lcd(0x27, 16, 2); // setup I2C for display

void setup() {
  lcd.init();
  lcd.backlight();
  
  last_time = millis();
  last_update = last_time;
  last_pot = digitalRead(POT_PIN);

  pinMode(POT_PIN,   INPUT);
  pinMode(GAS_PIN,   INPUT);
  pinMode(BRAKE_PIN, INPUT);
}

void loop() {

  read_values();
  calc_deltatime(); 
  calc_speed();
  print_display();
  
}

