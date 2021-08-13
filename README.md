# Face Recognition Clock In & Out

Using OpenVINO and Raspberry Pi 4

<img src="https://github.com/rainyyyan/face_recognition_clock_in_out/blob/master/images/excalidraw.png?raw=true78-ddfe8c3412a7.png" width=50% height=50%>

### Libraries Used

Python: 

- Requests

Go:

- GORM
- Beego
- Gin
- Viper

### To Run

python3 ./face_recognition_demo.py -i path/to/video -m_fd path/to/face_detection_model -m_lm path/to/landmarks_detection_model -m_reid path/to/face_reidentification_model -d_fd MYRIAD -d_lm MYRIAD -d_reid MYRIAD -fg path/to/face_gallery --no_show
