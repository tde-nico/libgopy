#ifndef LIBGOPY_H
# define LIBGOPY_H

# include <stdlib.h>

typedef __int64_t i64;
typedef __uint64_t u64;

typedef enum e_type
{
	TYPE_INT,
	TYPE_UINT,
	TYPE_FLOAT,
	TYPE_BYTES
}	t_type;

typedef struct s_pybytes
{
	unsigned char	*bytes;
	unsigned int	size;
}	t_pybytes;

typedef struct s_pyargs
{
	void			*value;
	t_type			t;
	struct s_pyargs	*next;
}	t_pyargs;

# ifdef __cplusplus
extern "C" {
# endif

void				init(void);
int					load(const char *module);
void				finalize(void);

void		call(const char *fname, int count, t_pyargs *args);
double		call_f64(const char *fname, int count, t_pyargs *args);
i64			call_i64(const char *fname, int count, t_pyargs *args);
u64			call_u64(const char *fname, int count, t_pyargs *args);
t_pybytes	call_byte(const char *fname, int count, t_pyargs *args);


# ifdef __cplusplus
}
# endif

#endif