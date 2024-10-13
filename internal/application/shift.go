package application

import "sort"

func (app *Application) generateShift(availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail) [5][7][]string {
	shift := [5][7][]string{}
	app.schedule1and2Slot(availableAssistants, assistantsDetailMap, &shift)
	app.schedule3Slot(availableAssistants, assistantsDetailMap, &shift)
	app.schedule4Slot(availableAssistants, assistantsDetailMap, &shift)
	app.schedule5Slot(availableAssistants, assistantsDetailMap, &shift)

	return shift
}

func (app *Application) schedule1and2Slot(availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string) {
	// 只有周一到周六需要安排 09:00-10:00 和 10:00-12:00 的值班
	for day := 1; day <= 6; day++ {
		app.schedule1and2SlotHelper(day, availableAssistants, assistantsDetailMap, shift, true, true)
		app.schedule1and2SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
		app.schedule1and2SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
		app.schedule1and2SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
	}
}

func (app *Application) schedule1and2SlotHelper(day int, availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string, requiredLeader bool, requiredSenior bool) {
	// 先看看有没有助理只填了 09:00-10:00
	assistantsOnly4Slot1 := []string{}
	for _, assistantInSlot1 := range availableAssistants[0][day] {
		flag := true
		for _, assistantInSlot2 := range availableAssistants[1][day] {
			if assistantInSlot1 == assistantInSlot2 {
				flag = false
				break
			}
		}
		if flag {
			assistantsOnly4Slot1 = append(assistantsOnly4Slot1, assistantInSlot1)
		}
	}

	if len(assistantsOnly4Slot1) > 0 {
		assistant := app.getProperAssistant(assistantsOnly4Slot1, requiredLeader, requiredSenior, assistantsDetailMap)
		if assistant != "" {
			shift[0][day] = append(shift[0][day], assistant)

			detail := assistantsDetailMap[assistant]
			detail.assignedTimeSlotCnt++
			assistantsDetailMap[assistant] = detail

			availableAssistants[0][day] = removeAssistant(availableAssistants[0][day], assistant)
		}
	}

	assistant := app.getProperAssistant(availableAssistants[1][day], requiredLeader, requiredSenior, assistantsDetailMap)
	if assistant != "" {
		if len(shift[0][day]) == len(shift[1][day]) {
			shift[0][day] = append(shift[0][day], assistant)

			detail := assistantsDetailMap[assistant]
			detail.assignedTimeSlotCnt++
			assistantsDetailMap[assistant] = detail

			availableAssistants[0][day] = removeAssistant(availableAssistants[0][day], assistant)
		}

		shift[1][day] = append(shift[1][day], assistant)

		detail := assistantsDetailMap[assistant]
		detail.assignedTimeSlotCnt++
		assistantsDetailMap[assistant] = detail

		availableAssistants[1][day] = removeAssistant(availableAssistants[1][day], assistant)
	} else {
		if len(shift[0][day]) == len(shift[1][day]) {
			shift[0][day] = append(shift[0][day], "")
		}
		shift[1][day] = append(shift[1][day], "")
	}
}

func (app *Application) getProperAssistant(assistants []string, requireLeader bool, requireSenior bool, assistantsDetailMap map[string]Detail) string {
	var limit = 1

	// 先按照可用时间段数排序
	sort.SliceStable(assistants, func(i, j int) bool {
		return assistantsDetailMap[assistants[i]].availableTimeSlotCnt < assistantsDetailMap[assistants[j]].availableTimeSlotCnt
	})

	// 如果需要负责人，那么先找负责人
	if requireLeader {
		var fallbackLeader string
		minAssignedTimeSlotCnt := int(^uint(0) >> 1) // max int

		for _, assistant := range assistants {
			if assistantsDetailMap[assistant].isEligibleAsLeader {
				if assistantsDetailMap[assistant].assignedTimeSlotCnt <= limit {
					return assistant
				}
				if assistantsDetailMap[assistant].assignedTimeSlotCnt < minAssignedTimeSlotCnt {
					minAssignedTimeSlotCnt = assistantsDetailMap[assistant].assignedTimeSlotCnt
					fallbackLeader = assistant
				}
			}
		}

		return fallbackLeader
	}

	// 如果需要资深助理，那么先找资深助理
	if requireSenior {
		var fallbackSenior string
		minAssignedTimeSlotCnt := int(^uint(0) >> 1) // max int

		for _, assistant := range assistants {
			if !assistantsDetailMap[assistant].isNewAssistant {
				if assistantsDetailMap[assistant].assignedTimeSlotCnt <= limit {
					return assistant
				}
				if assistantsDetailMap[assistant].assignedTimeSlotCnt < minAssignedTimeSlotCnt {
					minAssignedTimeSlotCnt = assistantsDetailMap[assistant].assignedTimeSlotCnt
					fallbackSenior = assistant
				}
			}
		}

		return fallbackSenior
	}

	var fallbackAssistant string
	minAssignedTimeSlotCnt := int(^uint(0) >> 1) // max int
	for _, assistant := range assistants {
		if assistantsDetailMap[assistant].assignedTimeSlotCnt <= limit {
			return assistant
		}
		if assistantsDetailMap[assistant].assignedTimeSlotCnt < minAssignedTimeSlotCnt {
			minAssignedTimeSlotCnt = assistantsDetailMap[assistant].assignedTimeSlotCnt
			fallbackAssistant = assistant
		}
	}
	return fallbackAssistant
}

func removeAssistant(slice []string, assistant string) []string {
	for i, v := range slice {
		if v == assistant {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (app *Application) schedule3Slot(availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string) {
	// 只有周一到周五需要安排 13:30-16:10 的值班
	for day := 1; day <= 5; day++ {
		app.schedule3SlotHelper(day, availableAssistants, assistantsDetailMap, shift, true, true)
		app.schedule3SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
		app.schedule3SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
		app.schedule3SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
	}
}

func (app *Application) schedule3SlotHelper(day int, availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string, requiredLeader bool, requiredSenior bool) {
	assistant := app.getProperAssistant(availableAssistants[2][day], requiredLeader, requiredSenior, assistantsDetailMap)
	if assistant != "" {
		shift[2][day] = append(shift[2][day], assistant)

		detail := assistantsDetailMap[assistant]
		detail.assignedTimeSlotCnt++
		assistantsDetailMap[assistant] = detail

		availableAssistants[2][day] = removeAssistant(availableAssistants[2][day], assistant)
	} else {
		shift[2][day] = append(shift[2][day], "")
	}
}

func (app *Application) schedule4Slot(availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string) {
	// 只有周一到周五需要安排 16:10-18:00 的值班
	for day := 1; day <= 5; day++ {
		app.schedule4SlotHelper(day, availableAssistants, assistantsDetailMap, shift, true, true)
		app.schedule4SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
		app.schedule4SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
		app.schedule4SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
	}
}

func (app *Application) schedule4SlotHelper(day int, availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string, requiredLeader bool, requiredSenior bool) {
	assistant := app.getProperAssistant(availableAssistants[3][day], requiredLeader, requiredSenior, assistantsDetailMap)
	if assistant != "" {
		shift[3][day] = append(shift[3][day], assistant)

		detail := assistantsDetailMap[assistant]
		detail.assignedTimeSlotCnt++
		assistantsDetailMap[assistant] = detail

		availableAssistants[3][day] = removeAssistant(availableAssistants[3][day], assistant)
	} else {
		shift[3][day] = append(shift[3][day], "")
	}
}

func (app *Application) schedule5Slot(availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string) {
	// 每天都有 19:00-21:00 的值班
	for day := 0; day <= 6; day++ {
		app.schedule5SlotHelper(day, availableAssistants, assistantsDetailMap, shift, true, true)
		app.schedule5SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
		app.schedule5SlotHelper(day, availableAssistants, assistantsDetailMap, shift, false, false)
	}
}

func (app *Application) schedule5SlotHelper(day int, availableAssistants [5][7][]string, assistantsDetailMap map[string]Detail, shift *[5][7][]string, requiredLeader bool, requiredSenior bool) {
	assistant := app.getProperAssistant(availableAssistants[4][day], requiredLeader, requiredSenior, assistantsDetailMap)
	if assistant != "" {
		shift[4][day] = append(shift[4][day], assistant)

		detail := assistantsDetailMap[assistant]
		detail.assignedTimeSlotCnt++
		assistantsDetailMap[assistant] = detail

		availableAssistants[4][day] = removeAssistant(availableAssistants[4][day], assistant)
	} else {
		shift[4][day] = append(shift[4][day], "")
	}
}
