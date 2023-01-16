#include <stdio.h>
#include <time.h>

#define year 1905
#define firstAdvent 27
#define secondAdvent 7
#define thirdAdvent 14
#define fourthAdvent 21
#define hours 0
#define minutes 0
#define seconds 0

void calculateAdventSundays() {
    // This initializes a struct tm (a structure representing a calendar date and time broken down into its components) called "first".
    struct tm first;
        // tm_year is defined as years since 1900
        first.tm_year = year - 1900;
        // month of the year [0,11]
        first.tm_mon = 10; // November
        first.tm_mday = firstAdvent;
        first.tm_hour = hours; first.tm_min = minutes; first.tm_sec = seconds;
    mktime(&first);

    // The while loop will keep incrementing the day of the month until the day of the week becomes Sunday(0).
    while(first.tm_wday != 0) {
        first.tm_mday++;
        mktime(&first);
    }
    
    char buffer[26];
    strftime(buffer, 26, "%Y-%m-%d %H:%M:%S", &first);
    printf("The first Sunday of Advent in %d is on %s\n", year, buffer);

    // Calculate the dates for the remaining Advent Sundays
    struct tm second = first;
    second.tm_mday += secondAdvent;
    mktime(&second);
    strftime(buffer, 26, "%Y-%m-%d %H:%M:%S", &second);
    printf("The second Sunday of Advent in %d is on %s\n", year, buffer);

    struct tm third = first;
    third.tm_mday += thirdAdvent;
    mktime(&third);
    strftime(buffer, 26, "%Y-%m-%d %H:%M:%S", &third);
    printf("The third Sunday of Advent in %d is on %s\n", year, buffer);

    struct tm fourth = first;
    fourth.tm_mday += fourthAdvent;
    mktime(&fourth);
    strftime(buffer, 26, "%Y-%m-%d %H:%M:%S", &fourth);
    printf("The fourth Sunday of Advent in %d is on %s\n", year, buffer);
}

int main(void) {

    printf("# C Snippets\n");
    printf("## Advent\n\n");

    calculateAdventSundays();

    return 0;
}

