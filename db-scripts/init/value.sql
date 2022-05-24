
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
	VALUES (59,1,'Tin tức công ty','Tin tức công ty 12/2021 (hynews)',NULL,NULL,current_timestamp),
(63,1,'Sự kiện nội bộ','Sự kiện nội bộ 12/2021',NULL,NULL,current_timestamp),
(64,1,'Thông báo','Thông báo tháng 12/2021',NULL,NULL,current_timestamp),
(65,2,'Chân dung nhân vật','Chân dung nhân vật tháng 12/2021',NULL,NULL,current_timestamp),
(66,2,'Chúc mừng sinh nhật','Chúc mừng sinh nhật các nhân viên sinh tháng 12/2021',NULL,NULL,current_timestamp),
(67,2,'Tâm tình','Tâm tình đồng nghiệp',NULL,NULL,current_timestamp),
(68,2,'Q&A','1000 câu hỏi vì sao?',NULL,NULL,current_timestamp),
(69,3,'E-learning','E-learning  tháng 12/2021',NULL,NULL,current_timestamp),
(70,3,'Tech corner','Tech corner tháng 12/2021',NULL,NULL,current_timestamp),
(60,2,'Dự án mới','Dự án mới 12/2021',NULL,NULL,current_timestamp),
(61,1,'Tin tức công ty 2022','Tin tức công ty 2022',NULL,NULL,current_timestamp),
(62,1,'Tin tức công ty 3','Tin tức công ty 3',NULL,'Tin tức công ty',current_timestamp),
(72,1,'Dự án mới','Dự án mới tháng 12/2021',NULL,'Tin tức công ty',current_timestamp),
(73,1,'Tiến độ/ Release dự án','tháng 12/2021',NULL,'Tin tức công ty',current_timestamp),
(74,1,'Thành tựu nghiên cứu','Thành tựu nghiên cứu năm 2021',NULL,'Tin tức công ty',current_timestamp),
(75,1,'Tin hợp tác','Tin hợp tác cuối năm 2021',NULL,'Tin tức công ty',current_timestamp),
(76,1,'Thông tin khác về công ty','những bí mật thú vị về công ty Hybrid Technologies',NULL,'Tin tức công ty',current_timestamp),
(77,1,'Điểm tin sự kiện/ Game nội bộ/ ăn chơi 3 miền','Điểm tin sự kiện/ Game nội bộ/ ăn chơi 3 miền cuối năm 2021',NULL,'Sự kiện nội bộ',current_timestamp),
(78,1,'Lịch Game/ Event của IC','Lịch Game/ Event của IC năm 2022',NULL,'Sự kiện nội bộ',current_timestamp),
(79,1,'Game giải trí','Game giải trí lmht,đế chế,...',NULL,'Sự kiện nội bộ',current_timestamp),
(80,1,'Hoạt động nội bộ từng miền','Hoạt động nội bộ từng miền bắc',NULL,'Sự kiện nội bộ',current_timestamp),
(81,1,'Chính sách/ chế độ cập nhật','Chính sách/ chế độ cập nhật 12/2021',NULL,'Thông báo',current_timestamp),
(82,1,'Thông báo nhắc nhở từ IT/ ISMS','Thông báo nhắc nhở từ IT/ ISMS cuối năm 2021',NULL,'Thông báo',current_timestamp),
(83,2,'Phỏng vấn nhân vật','Phỏng vấn nhân vật nổi bậc năm 2021',NULL,'Chân dung nhân vật',current_timestamp),
(84,2,'Chân dung nhân viên mới','Chân dung nhân viên mới tháng 10',NULL,'Chân dung nhân vật',current_timestamp),
(85,2,'Sinh nhật member theo tháng','Sinh nhật tháng 12 ',NULL,'Chúc mừng sinh nhật',current_timestamp),
(86,2,'Radio nội bộ','Radio nội bộ hybrid',NULL,'Tâm tình',current_timestamp),
(87,2,'Confession','thích thì phải nói nhưu đói là phải ăn',NULL,'Tâm tình',current_timestamp),
(88,2,'Câu hỏi cho BLĐ + giải đáp','hỏi nhanh đáp gọn',68,'Q&A',current_timestamp),
(89,2,'Câu hỏi Hytalk','Câu hỏi Hytalk 2022',68,'Q&A',current_timestamp),
(90,3,'Cổng thông tin bài học của công ty','Cổng thông tin bài học của công ty hybrid 2022',69,'E-learning',current_timestamp),
(91,3,'Lịch đào tạo tháng tiếp theo Training Center','12/2021-3/2022',69,'E-learning',current_timestamp);            