# AutoFSA

Обновления
=
## Версия 2.5 (28.04.2023)
	[*] Изменено именование файлов
	[+] Для файла csv формат даты как 20.04.2023, так и 20-02-23 для совместимости.
## Версия 2.4 (25.04.2023)
	[+] Добавлена поддержка прямого экспорта из xlsx
	[+] Добавлено ведение лога преобразований в папке программы
	[*] Исправлена ошибка обработки входного файла
## Версия 2.3 (22.04.2023)
	[*]  Изменен формат даты для csv файла (для удобства экспорта из Excel): 
		было 2016-01-15, стало 15.01.2016	 
##  Версия 2.2 (12.04.2024)
	[*] Исправлена ошибка с потеряным нулем в СНИЛС
##  Версия 2.1 (12.04.2023)
	[+] Добавлена опция в сsv2xml для выбора Типа сохранения данных (TypeSaveMethod)
	[*] Исправлен пример file.csv

##  Версия 2.0 (11.04.2023)
	Релиз
----
Требования
=
~~~~~
 MS Windows 7/8/10+ (32 & 64 bit)
 Ubuntu (amd64)
 MacOS X 11.0+
~~~~~
Порядок работы
=
	1. Скачайте последний релиз и перейдите в папку AutoFSA
	2. Сохраните (экспортируйте из Excel) вашу таблицу с поверками в файл file.csv или копируем оригинал в файл file.xlsx в эту папку
	3. Запустите сценарий start.cmd в зависимости от того, куда мы хотим отправить сведения (в черновики или на публикацию).
	4. В папке появится файл fsa_upload.xml	
	5. Загружаем его в ЛК ФГИС

ВАЖНО! 
=
	Файл fsa_upload.xml должен быть в кодировке UTF-8

Благодарности
=
Благодарю пользователя Комета за шаблон xml, и всех кто оставлял отзывы и замечания по работе!

Замечания пишите мне в телеграм https://t.me/makej4world (Демченко Артем Алекандрович).

	Апрель 2023