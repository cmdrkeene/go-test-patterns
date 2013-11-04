Describe("ScoreKeeper", func() {
    var scoreKeeper *ScoreKeeper

    BeforeEach(func() {
        scoreKeeper, err := NewScoreKeeper("Denver Broncos")
        Expect(err).NotTo(HaveOccured())
    })

    It("should have a starting score of 0", func() {
        Expect(scoreKeeper.Score).To(Equal(0))
    })

    Context("when a touchdown is scored", func() {
        BeforeEach(func() {
            scoreKeeper.Touchdown("Manning")
        })

        It("should increment the score", func() {
            Expect(scoreKeeper.Score).To(Equal(6))
        })
    })
})
