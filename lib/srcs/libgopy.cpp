#include "libgopy.h"
#include <iostream>
#include <Python.h>
#include <map>


std::map<std::string, PyObject *>	funcs;


void	init(void)
{
	Py_Initialize();
	PyRun_SimpleString("import sys; sys.path.insert(0, '.')");
}


int	load(const char *module)
{
	PyObject	*name;
	PyObject	*load_module;
	PyObject	*dir_list;
	Py_ssize_t	dir_size;
	PyObject	*attr_name;
	const char	*attr_str;
	PyObject	*func;

	name = PyUnicode_FromString(module);
	if (!name) {
		std::cerr << "Error: converting \"" << module << "\" into string\n";
		return (1);
	}

	load_module = PyImport_Import(name);
	if (!load_module) {
		std::cerr << "Error: importing \"" << module << "\" module\n";
		return (1);
	}

	dir_list = PyObject_Dir(load_module);
	if (!dir_list) {
		std::cerr << "Error: getting attributes of the module\n";
		return (1);
	}

	dir_size = PyList_Size(dir_list);
	for (Py_ssize_t i = 0; i < dir_size; i++) {
		attr_name = PyList_GetItem(dir_list, i);
		if (!attr_name) {
			std::cerr << "Error: getting attribute name\n";
			return (1);
		}

		attr_str = PyUnicode_AsUTF8(attr_name);
		if (!attr_str) {
			std::cerr << "Error: converting attribute name to string\n";
			return (1);
		}

		if (attr_str[0] == '_')
			continue;

		func = PyObject_GetAttrString(load_module, attr_str);
		if (!func) {
			std::cerr << "Error: finding \"" << func << "\" function\n";
			return (1);
		}

		funcs.insert(std::pair<std::string, PyObject *>(attr_str, func));
	}
	Py_DECREF(dir_list);

	return (0);
}

void	finalize(void)
{
	Py_Finalize();
}
