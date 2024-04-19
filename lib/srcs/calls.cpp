#include "libgopy.h"
#include <iostream>
#include <Python.h>


extern PyObject	*name;
extern PyObject	*load_module;

double	call_f64(const char *fname, int count, t_pyargs *args)
{
	PyObject	*func;
	PyObject	*res_obj;
	double		res;

	PyObject	*pyargs;
	PyObject	*arg;

	pyargs = PyTuple_New(count);
	int i = -1;
	while (++i < count && args)
	{
		switch (args->t) {
			case 'f':
				std::cout << "arg: " << *(double *)args->value << std::endl;
				arg = PyFloat_FromDouble(*(double *)args->value);
				break ;
			default:
				std::cerr << "Error: unknown type\"" << args->t << "\"\n";
				return (0);
		}
		PyTuple_SetItem(pyargs, i, arg);
		args = args->next;
	}

	if (!load_module) {
		std::cerr << "Error: no module imported\n";
		return (0);
	}
	func = PyObject_GetAttrString(load_module, fname);
	if (!func) {
		std::cerr << "Error: finding \"" << func << "\" function\n";
		return (0);
	}
	res_obj = PyObject_CallObject(func, pyargs);
	if (!res_obj) {
		std::cerr << "Error: calling \"" << func << "\" function\n";
		return (0);
	}
	res = PyFloat_AsDouble(res_obj);
	return (res);
}

long	call_i64(const char *fname)
{
	PyObject	*func;
	PyObject	*res_obj;
	long		res;

	if (!load_module) {
		std::cerr << "Error: no module imported\n";
		return (0);
	}
	func = PyObject_GetAttrString(load_module, fname);
	if (!func) {
		std::cerr << "Error: finding \"" << func << "\" function\n";
		return (0);
	}
	res_obj = PyObject_CallObject(func, NULL);
	if (!res_obj) {
		std::cerr << "Error: calling \"" << func << "\" function\n";
		return (0);
	}
	res = PyLong_AsLong(res_obj);
	return (res);
}

t_pybytes	call_byte(const char *fname)
{
	PyObject	*func;
	PyObject	*res_obj;
	t_pybytes	res;

	res.bytes = NULL;
	res.size = 0;
	if (!load_module) {
		std::cerr << "Error: no module imported\n";
		return (res);
	}
	func = PyObject_GetAttrString(load_module, fname);
	if (!func) {
		std::cerr << "Error: finding \"" << func << "\" function\n";
		return (res);
	}
	res_obj = PyObject_CallObject(func, NULL);
	if (!res_obj) {
		std::cerr << "Error: calling \"" << func << "\" function\n";
		return (res);
	}
	res.bytes = (unsigned char *)PyBytes_AsString(res_obj);
	res.size = PyBytes_Size(res_obj);
	return (res);
}
