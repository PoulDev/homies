package models

type Avatar struct {
	BgColor   string  `json:"bg_color"`
	FaceColor string  `json:"face_color"`
	FaceX     float32 `json:"face_x"`
	FaceY     float32 `json:"face_y"`
	LeX       float32 `json:"le_x"`
	LeY       float32 `json:"le_y"`
	ReX       float32 `json:"re_x"`
	ReY       float32 `json:"re_y"`
	Bezier    string  `json:"bezier"`
}
