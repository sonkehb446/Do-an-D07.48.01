
INSERT INTO public.branches(id, name)
	VALUES (1,'HTV-HAN'),(2,'SGN'),(3,'DAD'),(4,'HTJ');


INSERT INTO public.menus(id, name)
	VALUES (1,'HYNEWS'),(2,'HYBRIDERS'),(3,'HYLEARNING');

INSERT INTO public.positions(id, name)
	VALUES (1,'Fresher'),(2,'Tester'),
		 	(3,'Leader'),(4,'Manager'),
			(5,'Manager'),(6,'HR Manager');

INSERT INTO public.departments(id, name)
	VALUES (1,'IT'),(2,'HR'),(3,'accounting'),(4,'Web design'),(5,'IC');            


INSERT INTO public.images(id, url)
	VALUES (53,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/22f992a1-63a4-11ec-87c4-18dbf2313a75.png'),(26,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/bf879c8b-634d-11ec-a1e7-00090faa0001.jpg'),
		 	(27,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/profile/4672b2bd-6399-11ec-b7e4-18dbf2313a75.jpg'),(28,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/5f23b71e-639b-11ec-87c4-18dbf2313a75.jpg'),
			(29,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/e386fc27-639b-11ec-87c4-18dbf2313a75.jpg'),(30,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/4593e029-639c-11ec-87c4-18dbf2313a75.jpg'),
			(31,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/ee24a62b-639c-11ec-87c4-18dbf2313a75.jpg'),(32,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/224f9533-639d-11ec-87c4-18dbf2313a75.jpg'),
			(33,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/86630518-639d-11ec-87c4-18dbf2313a75.jpg'),
			(34,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/2ec550f5-639e-11ec-87c4-18dbf2313a75.jpg'),
			(35,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/921b0fcc-639e-11ec-87c4-18dbf2313a75.png'),
			(36,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/f5ee4234-639e-11ec-87c4-18dbf2313a75.jpg'),
			(37,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/60fe3ad7-639f-11ec-87c4-18dbf2313a75.jpg'),
			(38,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/9a2466d8-639f-11ec-87c4-18dbf2313a75.png'),
			(39,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/ce3576d6-639f-11ec-87c4-18dbf2313a75.jpg'),
			(40,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/feb91d76-639f-11ec-87c4-18dbf2313a75.png'),
			(41,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/98dd8f03-63a0-11ec-87c4-18dbf2313a75.jpg'),
			(42,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/36ee0924-63a1-11ec-87c4-18dbf2313a75.jpg'),
			(43,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/post/9d976fd6-63a1-11ec-87c4-18dbf2313a75.png'),
			(44,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/a864cb85-63a1-11ec-87c4-18dbf2313a75.png'),
			(45,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/dabc7ecf-63a1-11ec-87c4-18dbf2313a75.png'),
			(46,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/0f947a07-63a2-11ec-87c4-18dbf2313a75.png'),
			(47,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/56d08829-63a2-11ec-87c4-18dbf2313a75.png'),
			(48,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/91f91215-63a2-11ec-87c4-18dbf2313a75.png'),
			(49,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/ce2b4af6-63a2-11ec-87c4-18dbf2313a75.png'),
			(50,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/20fa4c07-63a3-11ec-87c4-18dbf2313a75.png'),
			(51,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/621b2598-63a3-11ec-87c4-18dbf2313a75.png'),
			(52,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/b46d713d-63a3-11ec-87c4-18dbf2313a75.jpg'),
			(54,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/5ac1777d-63a4-11ec-87c4-18dbf2313a75.png'),
			(55,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/9fc071ab-63a4-11ec-87c4-18dbf2313a75.png'),
			(56,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/fea5a55d-63a4-11ec-87c4-18dbf2313a75.jpg'),
			(57,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/41d226ba-63a5-11ec-87c4-18dbf2313a75.jpg'),
           (1,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/profile/4672b2bd-6399-11ec-b7e4-18dbf2313a75.jpg'),
			(58,'https://hylife-dev.s3-ap-southeast-1.amazonaws.com/smallImage/24225575-63b1-11ec-8f69-18dbf2313a75.png');                    


INSERT INTO public.users(
	id, email, password, role_user, branch_id, department_id, position_id, image_id, full_name,create_by)
	VALUES (9999,'sonkehb096@gmail.com','$2a$10$OCFri7nr7PhP5YNnzFJCnuWp9/fG54MBTvVo6TCRPTP2SwToZllCO','true', 2, 5, 2, 27, 'HongSon', 9999);    

INSERT INTO public.categories(
	id, menu_id, category_name,description, parent_id, parent_name, created_at)
	VALUES (59,1,'Tin t???c c??ng ty','Tin t???c c??ng ty 12/2021 (hynews)',NULL,NULL,current_timestamp),
(63,1,'S??? ki???n n???i b???','S??? ki???n n???i b??? 12/2021',NULL,NULL,current_timestamp),
(64,1,'Th??ng b??o','Th??ng b??o th??ng 12/2021',NULL,NULL,current_timestamp),
(65,2,'Ch??n dung nh??n v???t','Ch??n dung nh??n v???t th??ng 12/2021',NULL,NULL,current_timestamp),
(66,2,'Ch??c m???ng sinh nh???t','Ch??c m???ng sinh nh???t c??c nh??n vi??n sinh th??ng 12/2021',NULL,NULL,current_timestamp),
(67,2,'T??m t??nh','T??m t??nh ?????ng nghi???p',NULL,NULL,current_timestamp),
(68,2,'Q&A','1000 c??u h???i v?? sao?',NULL,NULL,current_timestamp),
(69,3,'E-learning','E-learning  th??ng 12/2021',NULL,NULL,current_timestamp),
(70,3,'Tech corner','Tech corner th??ng 12/2021',NULL,NULL,current_timestamp),
(60,2,'D??? ??n m???i','D??? ??n m???i 12/2021',NULL,NULL,current_timestamp),
(61,1,'Tin t???c c??ng ty 2022','Tin t???c c??ng ty 2022',NULL,NULL,current_timestamp),
(62,1,'Tin t???c c??ng ty 3','Tin t???c c??ng ty 3',NULL,'Tin t???c c??ng ty',current_timestamp),
(72,1,'D??? ??n m???i','D??? ??n m???i th??ng 12/2021',NULL,'Tin t???c c??ng ty',current_timestamp),
(73,1,'Ti???n ?????/ Release d??? ??n','th??ng 12/2021',NULL,'Tin t???c c??ng ty',current_timestamp),
(74,1,'Th??nh t???u nghi??n c???u','Th??nh t???u nghi??n c???u n??m 2021',NULL,'Tin t???c c??ng ty',current_timestamp),
(75,1,'Tin h???p t??c','Tin h???p t??c cu???i n??m 2021',NULL,'Tin t???c c??ng ty',current_timestamp),
(76,1,'Th??ng tin kh??c v??? c??ng ty','nh???ng b?? m???t th?? v??? v??? c??ng ty Hybrid Technologies',NULL,'Tin t???c c??ng ty',current_timestamp),
(77,1,'??i???m tin s??? ki???n/ Game n???i b???/ ??n ch??i 3 mi???n','??i???m tin s??? ki???n/ Game n???i b???/ ??n ch??i 3 mi???n cu???i n??m 2021',NULL,'S??? ki???n n???i b???',current_timestamp),
(78,1,'L???ch Game/ Event c???a IC','L???ch Game/ Event c???a IC n??m 2022',NULL,'S??? ki???n n???i b???',current_timestamp),
(79,1,'Game gi???i tr??','Game gi???i tr?? lmht,????? ch???,...',NULL,'S??? ki???n n???i b???',current_timestamp),
(80,1,'Ho???t ?????ng n???i b??? t???ng mi???n','Ho???t ?????ng n???i b??? t???ng mi???n b???c',NULL,'S??? ki???n n???i b???',current_timestamp),
(81,1,'Ch??nh s??ch/ ch??? ????? c???p nh???t','Ch??nh s??ch/ ch??? ????? c???p nh???t 12/2021',NULL,'Th??ng b??o',current_timestamp),
(82,1,'Th??ng b??o nh???c nh??? t??? IT/ ISMS','Th??ng b??o nh???c nh??? t??? IT/ ISMS cu???i n??m 2021',NULL,'Th??ng b??o',current_timestamp),
(83,2,'Ph???ng v???n nh??n v???t','Ph???ng v???n nh??n v???t n???i b???c n??m 2021',NULL,'Ch??n dung nh??n v???t',current_timestamp),
(84,2,'Ch??n dung nh??n vi??n m???i','Ch??n dung nh??n vi??n m???i th??ng 10',NULL,'Ch??n dung nh??n v???t',current_timestamp),
(85,2,'Sinh nh???t member theo th??ng','Sinh nh???t th??ng 12 ',NULL,'Ch??c m???ng sinh nh???t',current_timestamp),
(86,2,'Radio n???i b???','Radio n???i b??? hybrid',NULL,'T??m t??nh',current_timestamp),
(87,2,'Confession','th??ch th?? ph???i n??i nh??u ????i l?? ph???i ??n',NULL,'T??m t??nh',current_timestamp),
(88,2,'C??u h???i cho BL?? + gi???i ????p','h???i nhanh ????p g???n',68,'Q&A',current_timestamp),
(89,2,'C??u h???i Hytalk','C??u h???i Hytalk 2022',68,'Q&A',current_timestamp),
(90,3,'C???ng th??ng tin b??i h???c c???a c??ng ty','C???ng th??ng tin b??i h???c c???a c??ng ty hybrid 2022',69,'E-learning',current_timestamp),
(91,3,'L???ch ????o t???o th??ng ti???p theo Training Center','12/2021-3/2022',69,'E-learning',current_timestamp);            