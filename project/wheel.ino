void calc_deltatime() {
  float current_time = millis();
  deltatime = current_time - last_time;
  last_time = current_time;
}

void calc_speed() {
  float last_speed = speed; 
  if (gas_state && !break_state && speed < MAX_SPEED) { 
     speed += 0.1f * deltatime;
  } else if (break_state && speed > 0) {
    speed -= 0.1f * deltatime;
  }
  
  if (deltatime != 0.0f)
    acceleration = (speed - last_speed)/(deltatime / 1000);
}

void read_values(){
  pot_value = analogRead(POT_PIN);
  gas_state   = digitalRead(GAS_PIN);
  break_state = digitalRead(BRAKE_PIN);
}

void print_display() {
  float current_time = millis();
  if (current_time - last_update >= cool_down) { 
    lcd.clear();
    lcd.setCursor(1, 0);
    lcd.print("speed:");
    lcd.print(speed);
    lcd.setCursor(0, 1);
    last_update = current_time;
  }

}

