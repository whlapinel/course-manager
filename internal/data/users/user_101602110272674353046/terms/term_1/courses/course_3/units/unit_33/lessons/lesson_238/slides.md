---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

Write a function that accepts an integer and returns its cube

# **What is IoT Programming?**

**Definition**:  
IoT programming involves creating software to control and communicate with IoT devices, which are interconnected devices that collect, send, and act on data.

**Key Features**:

- Interaction with hardware sensors and actuators.
- Network communication between devices and servers.
- Real-time processing and cloud integration.

**Examples**:

- Smart home devices (e.g., thermostats, lights).
- Wearable fitness trackers.
- Industrial automation sensors.

# **Common IoT Programming Languages**

**Popular Languages for IoT**:

1. **Python**: Easy to use, with libraries like `RPi.GPIO` and `Adafruit_Blinka` for hardware interaction.
2. **C/C++**: Efficient and widely used in embedded systems (e.g., Arduino).
3. **JavaScript (Node.js)**: Great for web-connected IoT applications.
4. **Java**: Used for Android-based IoT devices.
5. **Rust**: Gaining popularity for its safety and performance.

# **Key Concepts in IoT Programming**

1. **Sensors and Actuators**:
   - Sensors collect data (e.g., temperature, light).
   - Actuators perform actions (e.g., turn on a motor).

2. **Communication Protocols**:
   - **MQTT**: Lightweight messaging protocol for IoT.
   - **HTTP/HTTPS**: For web-based IoT communication.
   - **Bluetooth/Wi-Fi**: For short-range communication.

3. **Edge Computing**:
   - Processing data locally on the device before sending it to the cloud.

# **Example Code (Python)**

**Task**: Control an LED with a button.

```python
import RPi.GPIO as GPIO
import time

# Set up GPIO pins
LED_PIN = 18
BUTTON_PIN = 23
GPIO.setmode(GPIO.BCM)
GPIO.setup(LED_PIN, GPIO.OUT)
GPIO.setup(BUTTON_PIN, GPIO.IN, pull_up_down=GPIO.PUD_UP)

# Main loop
try:
    while True:
        if GPIO.input(BUTTON_PIN) == GPIO.LOW:
            GPIO.output(LED_PIN, GPIO.HIGH)  # Turn on LED
        else:
            GPIO.output(LED_PIN, GPIO.LOW)  # Turn off LED
        time.sleep(0.1)
except KeyboardInterrupt:
    GPIO.cleanup()  # Cleanup GPIO on exit
```

# **Applications of IoT Programming**

1. **Smart Cities**:
   - Traffic monitoring and smart streetlights.

2. **Healthcare**:
   - Remote patient monitoring and wearable devices.

3. **Industrial IoT (IIoT)**:
   - Predictive maintenance and automated factories.

4. **Home Automation**:
   - Smart locks, lighting, and voice assistants.

# Agenda

- Set up your Finch: [Finch Setup](https://learn.birdbraintechnologies.com/finch/python/install/1-1)

- Assignment 7.1: Complete Finch Programming Lesson 1 and submit your code: [Lesson 1](https://learn.birdbraintechnologies.com/finch/python/program/lesson-1-moving-and-turning)
- [Browser IDE](https://brython.birdbraintechnologies.com/)
- [Programming Reference](https://learn.birdbraintechnologies.com/finch/python/?robot=finch&software=python&moduleslide=&pg=library&r=&&moduleslide2=)
