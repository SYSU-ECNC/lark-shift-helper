package application

import "fmt"

func (app *Application) getAvailableAssistant(submissions []*Submission) ([5][7][]string, error) {
	availableAssistant := [5][7][]string{}

	convertTimeSlot2number := func(timeSlot string) (int, error) {
		switch timeSlot {
		case "9：00-10：00":
			return 0, nil
		case "10：00-12：00":
			return 1, nil
		case "13：30-16：10":
			return 2, nil
		case "16：10-18：00":
			return 3, nil
		case "19：00-21：00":
			return 4, nil
		default:
			return -1, fmt.Errorf("invalid time slot %s", timeSlot)
		}
	}

	for _, submission := range submissions {
		for day, timeSlots := range submission.availableTime {
			for _, timeSlot := range timeSlots {
				timeSlotNumber, err := convertTimeSlot2number(timeSlot)
				if err != nil {
					return availableAssistant, err
				}

				availableAssistant[timeSlotNumber][day] = append(availableAssistant[timeSlotNumber][day], submission.name)
			}
		}
	}

	return availableAssistant, nil
}
