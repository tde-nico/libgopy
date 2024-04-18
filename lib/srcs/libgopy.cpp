#include "libgopy.h"
#include <iostream>
#include <Python.h>

void	f1()
{
	Py_Initialize();
	PyRun_SimpleString("import sys; sys.path.insert(0, '.')");

	PyObject *name, *load_module, *func, *callfunc, *args;
	
	name = PyUnicode_FromString((char *)"test");
	load_module = PyImport_Import(name);

	func = PyObject_GetAttrString(load_module, (char *)"func1");
	callfunc = PyObject_CallObject(func, NULL);
	double func1_out = PyFloat_AsDouble(callfunc);
	std::cout << "func1 output: " << func1_out << std::endl;

	func = PyObject_GetAttrString(load_module, (char *)"func2");
	args = PyTuple_Pack(1, PyFloat_FromDouble(1.0));
	callfunc = PyObject_CallObject(func, args);
	double func2_out = PyFloat_AsDouble(callfunc);
	std::cout << "func2 output: " << func2_out << std::endl;

	func = PyObject_GetAttrString(load_module, (char *)"func3");
	args = PyTuple_Pack(2, PyFloat_FromDouble(2.0), PyFloat_FromDouble(3.0));
	callfunc = PyObject_CallObject(func, args);
	double func3_out = PyFloat_AsDouble(callfunc);
	std::cout << "func3 output: " << func3_out << std::endl;

	func = PyObject_GetAttrString(load_module, (char *)"func4");
	args = PyTuple_Pack(1, PyUnicode_FromString((char *)"world"));
	callfunc = PyObject_CallObject(func, args);
	std::string func4_out = _PyUnicode_AsString(callfunc);
	std::cout << "func4 output: " << func4_out << std::endl;

	Py_Finalize();
}


PyObject	*name;
PyObject	*load_module;

void	init(void)
{
	Py_Initialize();
	PyRun_SimpleString("import sys; sys.path.insert(0, '.')");
	
	name = NULL;
	load_module = NULL;
}


int	load(const char *module)
{
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
	return (0);
}

void	finalize(void)
{
	Py_Finalize();
}
