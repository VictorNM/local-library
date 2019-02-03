insert into genres values
	('gen001', 'Fantasy'),
	('gen002', 'Science Fiction'),
	('gen003', 'French Poetry');
	
insert into authors values
	('aut001', 'Patrick', 'Rothfuss', '1973-06-06', null),
	('aut002', 'Ben', 'Bova', '1932-11-8', null),
	('aut003', 'Isaac', 'Asimov', '1920-01-02', '1992-04-06'),
	('aut004', 'Bob', 'Billings', null, null),
	('aut005', 'Jim', 'Jones', '1971-12-16', null);
	
insert into books values
	('bk0001', 'The Name of the Wind (The Kingkiller Chronicle, #1)', 'aut001', 'I have stolen princesses back from sleeping barrow kings. I burned down the town of Trebon. I have spent the night with Felurian and left with both my sanity and my life. I was expelled from the University at a younger age than most people are allowed in. I tread paths by moonlight that others fear to speak of during day. I have talked to Gods, loved women, and written songs that make the minstrels weep.', '9781473211896'),
	('bk0002', 'The Wise Man''s Fear (The Kingkiller Chronicle, #2)', 'aut001', 'Picking up the tale of Kvothe Kingkiller once again, we follow him into exile, into political intrigue, courtship, adventure, love and magic... and further along the path that has turned Kvothe, the mightiest magician of his age, a legend in his own time, into Kote, the unassuming pub landlord.', '9788401352836'),
	('bk0003', 'The Slow Regard of Silent Things (Kingkiller Chronicle)','aut001', 'Deep below the University, there is a dark place. Few people know of it: a broken web of ancient passageways and abandoned rooms. A young woman lives there, tucked among the sprawling tunnels of the Underthing, snug in the heart of this forgotten place.', '9780756411336'),
	('bk0004', 'Apes and Angels', 'aut002', 'Humankind headed out to the stars not for conquest, nor exploration, nor even for curiosity. Humans went to the stars in a desperate crusade to save intelligent life wherever they found it. A wave of death is spreading through the Milky Way galaxy, an expanding sphere of lethal gamma ...', '9780765379528'),
	('bk0005', 'Death Wave', 'aut002', 'In Ben Bova''s previous novel New Earth, Jordan Kell led the first human mission beyond the solar system. They discovered the ruins of an ancient alien civilization. But one alien AI survived, and it revealed to Jordan Kell that an explosion in the black hole at the heart of the Milky Way galaxy has created a wave of deadly radiation, expanding out from the core toward Earth. Unless the human race acts to save itself, all life on Earth will be wiped out...', '9780765379504');
	
insert into books_genres values
	('bk0001', 'gen001'),
	('bk0002', 'gen001'),
	('bk0003', 'gen001'),
	('bk0004', 'gen002'),
	('bk0005', 'gen002');
	
insert into book_instances values
	('bi0001', 'bk0001', 'Available', null),
	('bi0002', 'bk0002', 'Loaned', null),
	('bi0003', 'bk0003', 'Loaned', null),
	('bi0004', 'bk0004', 'Available', null),
	('bi0005', 'bk0004', 'Available', null),
	('bi0006', 'bk0004', 'Available', null),
	('bi0007', 'bk0005', 'Maintenance', null);