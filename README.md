# Vimaan OS
 
Educational. Hobby.  
Personal spare-time/boredom initiative.  

> [!NOTE]
> Vimāna or VayuYaan are translations for Airplane.

## Overview

RC Planes communicate in real time in near-direct-line-of-sight at frequency ranges between few hundred MHz to few GHz.  

*Higher* Frequency => *Smaller* Wavelength => *Smaller* Antennae, *More* Bandwidth and *Lower* Latency **BUT** *Greater* Signal Attenuation. Linear Polarization has more attenuation than Circular Polarization but needs lower (electrical) power for similar range.  

Advanced/Mission critical use cases (often) have (proprietary) softwares that enable various modes of operations such as LOITER and REAL TIME THROTTLE CONTROL which is enabled by sophisticated hardware that are often regulated. Visuals and instructions are communicated via Satcom uplink and downlink.

Explorations for lightweight interface for small aerial vehicles ([IoT]).  

### Applications

- Post calamity (hurricanes, flooding, etc) inspection and analysis for search and life support.
- Aerial water sprinkling over fields / plantations that are far away from water sources.
- Relieving foot patrols from regions of high human risk / conflict.

Most of the above are are non commercial and non scalable, and do not incentivize development.

## Control By Command

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
