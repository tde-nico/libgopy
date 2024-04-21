#include "libgopy.h"
#include <iostream>
#include <Python.h>
#include <map>


extern std::map<std::string, PyObject *>	funcs;


template <typename K, typename V>
V GetWithDef(const std::map <K,V> &m, const K &key, const V &defval) {
	typename std::map<K,V>::const_iterator it = m.find(key);
	if (it == m.end()) {
		return (defval);
	} else {
		return (it->second);
	}
}


PyObject	*setup_args(int count, t_pyargs *args)
{
	PyObject	*pyargs;
	PyObject	*arg;

	if (!count)
		return (NULL);
	pyargs = PyTuple_New(count);
	int i = count;
	while (--i >= 0 && args)
	{
		switch (args->t) {
			case 'f':
				arg = PyFloat_FromDouble(*(double *)args->value);
				break ;
			case 'd':
				arg = PyLong_FromLong(*(int *)args->value);
				break ;
			case 'b':
				arg = PyBytes_FromString((const char *)args->value);
				break ;
			default:
				std::cerr << "Error: unknown type\"" << args->t << "\"\n";
				return (0);
		}
		if (PyTuple_SetItem(pyargs, i, arg))
			return (NULL);
		args = args->next;
	}
	return (pyargs);
}

PyObject	*call_func(const char *fname, int count, t_pyargs *args)
{
	PyObject	*func;
	PyObject	*res_obj;
	PyObject	*pyargs;

	func = NULL;
	func = GetWithDef(funcs, std::string(fname), func);
	if (!func) {
		std::cerr << "Error: finding \"" << fname << "\" function\n";
		return (NULL);
	}
	pyargs = setup_args(count, args);
	res_obj = PyObject_CallObject(func, pyargs);
	if (!res_obj) {
		std::cerr << "Error: calling \"" << func << "\" function\n";
		return (NULL);
	}
	return (res_obj);
}

void	call(const char *fname, int count, t_pyargs *args)
{
	call_func(fname, count, args);
}

double	call_f64(const char *fname, int count, t_pyargs *args)
{
	PyObject	*res_obj;
	double		res;

	res_obj = call_func(fname, count, args);
	if (!res_obj)
		return(0);
	res = PyFloat_AsDouble(res_obj);
	return (res);
}

long	call_i64(const char *fname, int count, t_pyargs *args)
{
	PyObject	*res_obj;
	long		res;

	res_obj = call_func(fname, count, args);
	if (!res_obj)
		return(0);
	res = PyLong_AsLong(res_obj);
	return (res);
}

t_pybytes	call_byte(const char *fname, int count, t_pyargs *args)
{
	PyObject	*res_obj;
	t_pybytes	res;

	res.bytes = NULL;
	res.size = 0;
	res_obj = call_func(fname, count, args);
	if (!res_obj)
		return(res);
	res.bytes = (unsigned char *)PyBytes_AsString(res_obj);
	res.size = PyBytes_Size(res_obj);
	return (res);
}
