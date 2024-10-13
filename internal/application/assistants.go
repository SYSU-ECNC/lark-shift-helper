package application

type Detail struct {
	isAvailableForFieldWork bool
	isEligibleAsLeader      bool
	isNewAssistant          bool
	availableTimeSlotCnt    int
	assignedTimeSlotCnt     int
}

func (app *Application) getAssistantsDetailMapFromSubmissions(submissions []*Submission) (map[string]Detail, error) {
	assistantsDetailMap := map[string]Detail{}

	getTimeSlotCntFromAvailableTime := func(availableTime [7][]string) int {
		cnt := 0
		for _, timeSlots := range availableTime {
			cnt += len(timeSlots)
		}
		return cnt
	}

	for _, submission := range submissions {
		assistantsDetailMap[submission.name] = Detail{
			isAvailableForFieldWork: submission.isAvailableForFieldWork,
			isEligibleAsLeader:      submission.isEligibleAsLeader,
			isNewAssistant:          submission.isNewAssistant,
			availableTimeSlotCnt:    getTimeSlotCntFromAvailableTime(submission.availableTime),
			assignedTimeSlotCnt:     0,
		}
	}

	return assistantsDetailMap, nil
}
