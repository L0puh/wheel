
void calc_deltatime() {
  float current_time = millis();
  deltatime = current_time - last_time;
  last_time = current_time;
}

void calc_speed() {
  float last_speed = speed; 
  if (gas_state && !break_state && speed < MAX_SPEED) { 
    speed += 0.01 * deltatime;
  } else if (break_state && speed > 0) {
    speed -= 0.01 * deltatime;
  }

  if (deltatime != 0.0f)
    acceleration = (speed - last_speed)/(deltatime / 1000);
}

void read_values(){
  pot_value = analogRead(POT_PIN);
  gas_state   = digitalRead(GAS_PIN);
  break_state = digitalRead(BRAKE_PIN);
}


void send_values(float speed, char dir, int factor) {

  /* 
   send data to utils's serial
   and process it in the game (simulation)
   TODO: 
   provide joystick functionality
   write a driver (?)
  */
  Serial.print("D:");
  Serial.print(dir);
  Serial.print(" S:");
  Serial.print(speed);
  Serial.print(" F:");
  Serial.println(factor);
}

change_t calc_change() {
  change_t ch;
  ch.factor = last_pot-pot_value;
  char dir; 

  /* 

     the value fluctuates without 
     touching the potentiometer in about
     1 unit, hence skip it

  */

  if (ch.factor > 1)
    dir =  'R';
  else if (ch.factor < -1) 
    dir = 'L';
  else 
    dir = 'N';

  ch.dir = dir;
  last_pot = pot_value;

  return ch;
}

void update() {
  float current_time = millis();
  change_t change = calc_change();
  if (current_time - last_update >= cool_down) { 
    lcd.clear();
    lcd.setCursor(1, 0);
    lcd.print("speed:");
    lcd.print(speed);
    lcd.setCursor(0, 1);
    lcd.print(change.dir);
    lcd.print(" ");
    lcd.print(change.factor);
    last_update = current_time;

  }
   send_values(speed, change.dir, change.factor);
}

