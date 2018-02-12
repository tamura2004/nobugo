cs = [
		["弥助", 1, 2, 6, "CHANGE_DICE"],
		["木下藤吉郎", 3, 4, 1, "CHANGE_DICE"],
		["明智光秀", 5, 6, 1, "CHANGE_DICE"],
		["柴田勝家", 3, 4, 5, "CHANGE_DICE"],
		["前田利家", 5, 6, 3, "CHANGE_DICE"],
		["滝川一益", 1, 2, 4, "CHANGE_DICE"],
		["毛利元就", 5, 6, 4, "CHANGE_DICE"],
		["徳川家康", 3, 4, 2, "CHANGE_DICE"],
		["今川義元", 5, 6, 1, "CHANGE_DICE"],
		["北条氏康", 1, 2, 3, "CHANGE_DICE"],
		["島津義弘", 3, 4, 6, "CHANGE_DICE"],
		["松永久秀", 0, 0, 0, "CHANGE_SAMURAI"],
		["武田信玄", 1, 0, 0, "ADD_GOLD"],
		["上杉謙信", 2, 0, 0, "ADD_BATTLE_DICE"],
		["伊達政宗", 2, 4, 0, "CHANGE_AREA"],
		["九鬼義隆", 2, 6, 0, "CHANGE_AREA"],
		["伊藤マンショ", 1, 6, 0, "CHANGE_AREA"],
		["足利義昭", 1, 0, 0, "DELETE_DICE"],
		["大友宗麟", 1, 4, 0, "CHANGE_AREA"],
]

cs.each do |c|
	df = []
	[1,2,3].each do |i|
		if c[i] != 0
			df << c[i]
		end
	end
	puts <<"EOD"
[[Samurai]]
Name = "#{c[0]}"
Ability = "#{c[4]}"
Power = #{df.inspect}

EOD
end
