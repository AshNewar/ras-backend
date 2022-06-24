package rc

import "github.com/gin-gonic/gin"

func addStudentResume(ctx *gin.Context, resume string, sid uint, rid uint) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Create(&StudentRecruitmentCycleResume{
		StudentRecruitmentCycleID: sid,
		Resume:                    resume,
		RecruitmentCycleID:        rid,
	})
	return tx.Error
}

func fetchStudentResume(ctx *gin.Context, sid uint, resumes *[]StudentRecruitmentCycleResume) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Where("student_recruitment_cycle_id = ?", sid).Find(resumes)
	return tx.Error
}

func fetchAllResumes(ctx *gin.Context, rid uint, resumes *[]AllResumeResponse) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).
		Joins("JOIN student_recruitment_cycles ON student_recruitment_cycles.id = student_recruitment_cycle_resumes.student_recruitment_cycle_id AND student_recruitment_cycle_resumes.recruitment_cycle_id = ?", rid).
		Select("student_recruitment_cycles.name as name, student_recruitment_cycles.email as email, student_recruitment_cycles.id as sid, student_recruitment_cycle_resumes.id as rsid, student_recruitment_cycle_resumes.resume as resume, student_recruitment_cycle_resumes.verified as verified, student_recruitment_cycle_resumes.action_taken_by as action_taken_by").
		Scan(resumes)
	return tx.Error
}

func fetchResume(ctx *gin.Context, rsid uint) (string, error) {
	var resume string
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Where("id = ?", rsid).Pluck("resume", &resume)
	return resume, tx.Error
}

func updateResumeVerify(ctx *gin.Context, rsid uint, verified bool, user string) (bool, error) {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Where("id = ?", rsid).Update("verified", verified).Update("action_taken_by", user)
	return tx.RowsAffected == 1, tx.Error
}
