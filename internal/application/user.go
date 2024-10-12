package application

type User struct {
	name                    string
	isAvailableForFieldWork bool
	isEligibleAsLeader      bool
	isNewAssistant          bool
	availableTimeSlotCnt    int
	assignedTimeSlotCnt     int
}

func (app *Application) getUsersFromSubmissions(submissions []*Submission) ([]*User, error) {
	users := []*User{}

	getTimeSlotCntFromAvailableTime := func(availableTime [7][]string) int {
		cnt := 0
		for _, timeSlots := range availableTime {
			cnt += len(timeSlots)
		}
		return cnt
	}

	for _, submission := range submissions {
		user := &User{
			name:                    submission.name,
			isAvailableForFieldWork: submission.isAvailableForFieldWork,
			isEligibleAsLeader:      submission.isEligibleAsLeader,
			isNewAssistant:          submission.isNewAssistant,
			availableTimeSlotCnt:    getTimeSlotCntFromAvailableTime(submission.availableTime),
			assignedTimeSlotCnt:     0,
		}

		users = append(users, user)
	}

	return users, nil
}
