local adjectives = "adjectives.txt"
local animals = "animals.txt"

function read_file_as_array(file_path)
	local file = io.open(file_path, "r")
	if not file then
		return nil, "Error Opening File"
	end

	-- table
	local lines = {}

	for line in file:lines() do
		table.insert(lines, line)
	end

	file:close()

	return lines
end

function get_name()
	local adjectiveTable = read_file_as_array(os.getenv("PWD") .. "/" .. adjectives)

	local animalTable = read_file_as_array(animals)

	local adjectiveId = math.floor(math.random() * (#adjectiveTable - 1) + 1)

	local animalId = math.floor(math.random() * (#animalTable - 1) + 1)

	local randomId = math.floor(math.random() * (1000 - 1) + 1)

	name = string.gsub(adjectiveTable[adjectiveId], " ", "-") .. "-" .. string.gsub(animalTable[animalId], " ", "-")

	return string.lower(name) .. "-" .. randomId
end

