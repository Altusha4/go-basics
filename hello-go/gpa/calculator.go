package gpa

func Calculate(percent int) (string, float64, string) {
	switch {
	case percent >= 95:
		return "A", 4.0, "Excellent"
	case percent >= 90:
		return "A-", 3.67, "Excellent"
	case percent >= 85:
		return "B+", 3.33, "Good"
	case percent >= 80:
		return "B", 3.0, "Good"
	case percent >= 75:
		return "B-", 2.67, "Good"
	case percent >= 70:
		return "C+", 2.33, "Satisfactory"
	case percent >= 65:
		return "C", 2.0, "Satisfactory"
	case percent >= 60:
		return "C-", 1.67, "Satisfactory"
	case percent >= 55:
		return "D+", 1.33, "Satisfactory"
	case percent >= 50:
		return "D", 1.0, "Satisfactory"
	case percent >= 25:
		return "FX", 0.0, "Fail"
	default:
		return "F", 0.0, "Fail"
	}
}
