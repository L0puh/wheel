#ifndef CONSTANTS_H
#define CONSTANTS_H 


#define POT_PIN   A0 // potentiometer
#define GAS_PIN   2  // button gas
#define BRAKE_PIN 3  // button slow down
#define MAX_SPEED 1000

typedef struct change_t {
  int factor = 0;
  char dir = 'N';
} change_t;

void update();
void calc_deltatime();
void calc_speed();
change_t calc_change();

void read_values();
void send_values(float, char, float);

#endif 
