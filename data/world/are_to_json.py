import json
import shlex

text = ''
last_str = ''

def read_char():
	global text
	global last_str
	text = text.lstrip()
	c = text[0]
	text = text[1:]
	last_str = c
	return c

def read_word():
	global text
	global last_str
	text = text.lstrip()

	if text[0] == "'":
		last_str, text = text[1:].split("'", 1)
		return "'" + last_str + "'"
	if text[0] == '"':
		last_str, text = text[1:].split('"', 1)
		return "'" + last_str + "'"

	buf = ''
	c = text[0]
	while c != ' ' and c != '\t' and c != '\n':
		buf += c
		text = text[1:]
		c = text[0]
	last_str = buf
	return buf
#	last_str, text = text.split(None, 1)
#	print last_str
#	return last_str.lstrip()

def read_string():
	global text
	global last_str
	last_str, text = text.split('~', 1)
	if last_str[0] == '\n':
		last_str = last_str[1:]
	return last_str

def read_int():
	global text
	global last_str
	buf = ''
	text = text.lstrip()
	if text[0] == '-' or text[0] == '+':
		buf = text[0]
		text = text[1:].lstrip()

	while text[0].isdigit():
		buf += text[0]
		text = text[1:]
	last_str = buf
	return int(buf)

def read_dice():
	global last_str
	buf = ''
	buf += '%d'%read_int()
	buf += read_char()
	buf += '%d'%read_int()
	buf += read_char()
	buf += '%d'%read_int()
	last_str = buf
	return buf

def read_eol():
	global text

	c = text[0]
	while c != '\n':
		text = text[1:]
		c = text[0]
#	text = text.split('\n', 1)[-1]

def read_area():
	area = {}
	area['filename'] = read_string()
	area['name'] = read_string()
	area['credits'] = read_string()
	area['min_vnum'] = read_int()
	area['max_vnum'] = read_int()
	return area

def read_mobile():
	global text
	text = text.lstrip()
	if text[0] != '#':
		print 'load_mobiles: # not found'
		exit()

	text = text[1:]
	vnum = read_int()

	if vnum == 0:
		return None

	mobile = {}
	mobile['vnum'] = vnum
	mobile['keywords'] = read_string()
	mobile['short_descr'] = read_string()
	mobile['long_descr'] = read_string()
	mobile['description'] = read_string()
	mobile['race'] = read_string()
	mobile['act_flags'] = read_word()
	mobile['aff_flags'] = read_word()
	mobile['alignment'] = read_int()
	mobile['group'] = read_word()
	mobile['level'] = read_int()
	mobile['hitroll'] = read_int()
	mobile['hp_dice'] = read_dice()
	mobile['mana_dice'] = read_dice()
	mobile['damage_dice'] = read_dice()
	mobile['damage_type'] = read_word()
	mobile['ac_pierce'] = read_int()
	mobile['ac_bash'] = read_int()
	mobile['ac_slash'] = read_int()
	mobile['ac_exotic'] = read_int()
	mobile['off_flags'] = read_word()
	mobile['imm_flags'] = read_word()
	mobile['res_flags'] = read_word()
	mobile['vuln_flags'] = read_word()
	mobile['start_pos'] = read_word()
	mobile['default_pos'] = read_word()
	mobile['sex'] = read_word()
	mobile['wealth'] = read_int()
	mobile['form_flags'] = read_word()
	mobile['part_flags'] = read_word()
	mobile['size'] = read_word()
	mobile['material'] = read_word()

	while True:
		text = text.lstrip()
		if text[0] == 'F':
			text = text[1:]
			if 'remove_flags' not in mobile:
				mobile['remove_flags'] = []
			flag_mod = {}
			flag_mod['type'] = read_word()
			flag_mod['bit'] = read_word()
			mobile['remove_flags'].append(flag_mod)
		elif text[0] == '>':
			text = text[1:]
			if 'mobprogs' not in mobile:
				mobile['mobprogs'] = []
			mobprog = {}
			parts = read_string().split(None, 1)
			mobprog['type'] = parts[0]
			if len(parts) > 1:
				mobprog['args'] = parts[1]
			mobprog['commands'] = read_string()
			mobile['mobprogs'].append(mobprog)
		elif text[0] == '|':
			text = text[1:]
			break
		elif text[0] == '#':
			break
		else:
			print 'weird mob', vnum
			exit()

	return mobile

def read_mobiles():
	mobiles = {}

	while True:
		mobile = read_mobile()
		if mobile == None:
			break
		mobiles[mobile['vnum']] = mobile

	return mobiles

def read_object():
	global text
	text = text.lstrip()
	if text[0] != '#':
		print 'load_object: # not found'
		exit()

	text = text[1:]
	vnum = read_int()

	if vnum == 0:
		return None

	obj = {}
	obj['vnum'] = vnum
	obj['keywords'] = read_string()
	obj['short_descr'] = read_string()
	obj['description'] = read_string()
	obj['material'] = read_string()
	obj['item_type'] = read_word()
	obj['extra_flags'] = read_word()
	obj['wear_flags'] = read_word()
	obj['value0'] = read_word()
	obj['value1'] = read_word()
	obj['value2'] = read_word()
	obj['value3'] = read_word()
	obj['value4'] = read_word()
	obj['level'] = read_int()
	obj['weight'] = read_int()
	obj['cost'] = read_int()
	obj['condition'] = read_word()

	while True:
		text = text.lstrip()
		if text[0] == 'E':
			text = text[1:]
			if 'extra_descr' not in obj:
				obj['extra_descr'] = {}
			obj['extra_descr'][read_string()] = read_string()
		elif text[0] == 'A':
			text = text[1:]
			if 'added_affects' not in obj:
				obj['added_affects'] = []
			affect = {}
			affect['location'] = read_int()
			affect['modifier'] = read_int()
			obj['added_affects'].append(affect)
		elif text[0] == 'F':
			text = text[1:]
			if 'added_flags' not in obj:
				obj['added_flags'] = []
			add = {}
			text = text.lstrip()
			add['type'] = text[0]
			text = text[1:]
			add['location'] = read_int()
			add['modifier'] = read_int()
			add['bitvector'] = read_word()
			obj['added_flags'].append(add)
		elif text[0] == '#':
			break
		else:
			print 'weird obj', vnum
			exit()

	return obj

def read_objects():
	objects = {}

	while True:
		obj = read_object()
		if obj == None:
			break
		objects[obj['vnum']] = obj

	return objects

def read_room():
	global text
	text = text.lstrip()
	if text[0] != '#':
		print 'load_room: # not found'
		exit()

	text = text[1:]
	vnum = read_int()

	if vnum == 0:
		return None

	room = {}
	room['vnum'] = vnum
	room['name'] = read_string()
	room['description'] = read_string()
	room['tele_dest'] = read_int()
	room['room_flags'] = read_word()
	room['sector_type'] = read_int()

	while True:
		c = read_char()
		if c == 'E':
			if 'extra_descr' not in room:
				room['extra_descr'] = {}
			room['extra_descr'][read_string()] = read_string()
		elif c == 'H':
			room['heal_rate'] = read_int()
		elif c == 'M':
			room['mana_rate'] = read_int()
		elif c == 'C':
			room['clan'] = read_word()
		elif c == 'G':
			room['guild'] = read_word()
		elif c == 'O':
			room['owner'] = read_string()
		elif c == 'D':
			if 'exits' not in room:
				room['exits'] = {}
			exit = {}
			direction = read_int()
			exit['description'] = read_string()
			exit['keywords'] = read_string()
			exit['locks'] = read_int()
			exit['key'] = read_int()
			exit['vnum'] = read_int()
			room['exits'][direction] = exit
		elif c == 'S':
			break
		else:
			print 'weird room', vnum
			exit()
	return room

def read_rooms():
	rooms = {}

	while True:
		room = read_room()
		if room == None:
			break
		rooms[room['vnum']] = room

	return rooms

def read_resets():
	global text
	resets = []

	while True:
		letter = read_char()

		if letter == '*':
			read_eol()
			continue
		elif letter == 'S':
			break

		reset = {}
		read_int()
		reset['type'] = letter
		reset['arg1'] = read_int()
		reset['arg2'] = read_int()
		if letter == 'G' or letter == 'R':
			reset['arg3'] = 0
		else:
			reset['arg3'] = read_int()
		if letter == 'P' or letter == 'M':
			reset['arg4'] = read_int()
		else:
			reset['arg4'] = 0

		read_eol()
		resets.append(reset)

	return resets

def read_shops():
	global text
	shops = []

	while True:
		vnum = read_int()

		if vnum == 0:
			break

		shop = {}
		shop['keeper'] = vnum
		shop['buy_type'] = read_int()
		shop['profit_buy'] = read_int()
		shop['profit_sell'] = read_int()
		shop['open_hour'] = read_int()
		shop['close_hour'] = read_int()

		read_eol()
		shops.append(shop)

	return shops

def read_specials():
	global text
	specials = []

	while True:
		letter = read_char()

		if letter == '*':
			read_eol()
			continue
		elif letter == 'S':
			break

		special = {}
		special['type'] = letter
		special['vnum'] = read_int()
		special['spec'] = read_word()

		read_eol()
#		print 'read special, remaining is', text
		specials.append(special)

#	print 'finished specials, remaining is', text
	return specials

def read_tourstarts():
	tourstarts = {}
	return tourstarts

def read_tourroutes():
	tourroutes = {}
	return tourroutes

def to_json(f):
	global text
	sections = {}

	for line in f:
		text += line

	text.replace(r'\r','')

	try:
		while True:
			text = text.lstrip()
			if (len(text) == 0):
				break;

			parts = text.split(None, 1)
			if len(parts) == 1:
#				print parts[0]
				break
			word = parts[0]
			text = parts[1]
	#		print word

			if word == '#AREA':
				sections['AREA'] = read_area()
				for k in sections['AREA']:
					sections[k] = sections['AREA'][k]
			elif word == '#MOBILES':
				sections['MOBILES'] = read_mobiles()
			elif word == '#OBJECTS':
				sections['OBJECTS'] = read_objects()
			elif word == '#RESETS':
				sections['RESETS'] = read_resets()
			elif word == '#ROOMS':
				sections['ROOMS'] = read_rooms()
			elif word == '#SHOPS':
				sections['SHOPS'] = read_shops()
			elif word == '#SPECIALS':
				sections['SPECIALS'] = read_specials()
	#		elif word == '#TOURSTARTS':
	#			sections['TOURSTARTS'] = read_tourstarts()
	#		elif word == '#TOURROUTES':
	#			sections['TOURROUTES'] = read_tourroutes()
			elif word.startswith('#$'):
				break
			else:
				pass
	except:
		print 'Last string:', last_str
		print 'Remaining:', text
		raise

#	return sections
	return json.dumps(sections, indent=4)

if __name__ == "__main__":
	import codecs
	import sys
#	area = open(sys.argv[1])
	area = codecs.open(sys.argv[1], 'r', 'utf-8', 'ignore')

	try:
		j = to_json(area)
	except:
		print sys.argv[1]
		raise
	finally:
		area.close()

	fname = sys.argv[1].rsplit('.', 1)[0]
	f = open(fname + '.json', 'w')
	f.write(j)
	f.close()
#	import pprint
#	pp = pprint.PrettyPrinter(depth=4)
#	pp.pprint(j)
#	print j

