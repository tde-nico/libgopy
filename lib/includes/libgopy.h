#ifndef LIBGOPY_H
# define LIBGOPY_H

typedef struct	s_pybytes
{
	unsigned char	*bytes;
	unsigned int	size;
}	t_pybytes;

typedef struct s_pyargs
{
	void			*value;
	char			t;
	struct s_pyargs	*next;
}	t_pyargs;

# ifdef __cplusplus
extern "C" {
# endif

void	init(void);
int		load(const char *module);
void	finalize(void);

double		call_f64(const char *fname, int count, t_pyargs *args);
long		call_i64(const char *fname);
t_pybytes	call_byte(const char *fname);


# ifdef __cplusplus
}
# endif

#endif