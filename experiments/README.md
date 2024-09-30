Experiments
===========

Python for model definition. Port to low level later.  

Notes
-----

SILABs CP210xxx for UART. Use the correct drivers.

Control By Command
------------------

A command could be as basic as setting the altitude to 1000 meters.

```
ALTITUDE 1000; // meters

```

This can be reduced to a lower level instruction set as a queue (FIFO execution), which can be further blown up into instructions that directly pertain to flight mechanics.

```
// FIFO Queue

PITCH 20; // degrees
WAIT UNTIL object.altitude == 1000; // meters
PITCH LEVEL;


// Further Blown Up

while (object.altitude < 1000 - delta) {
    object.elevator = elevator_control(object.pitch, 20);
    // Method signature
    // return_type elevator_control(current_pitch, required_pitch);

    // Maintain other parameters same as before
    // such as maintain speed same as before
    // (at 100 Kmph) say.
    object.throttle = throttle_control(object.speed, 100);
    // Method signature:
    // return_type throttle_control(current_speed, required_speed);
}

while (object.pitch != 0) {
    object.elevator = elevator_control(object.pitch, 0);
}

```

The control functions (like `throttle_control` and `elevator_control` above) can be deterministic equations or RL Agents.
