#####   COLORS   #####

END				= \033[0m

GREY			= \033[30m
RED				= \033[31m
GREEN			= \033[32m
YELLOW			= \033[33m
BLUE			= \033[34m
PURPLE			= \033[35m
CYAN			= \033[36m

HIGH_RED		= \033[91m

#####   RULES   #####

NAME	= test
LIB		= ../lib/libgopy.so

GB		= go build
#GOFLAGS	= -a -installsuffix cgo -ldflags '-s'

all:
	@ make -s -C ../lib
	@ cp $(LIB) .
	@ $(GB) $(GOFLAGS)
	@ echo "$(GREEN)[+] $(NAME)$(END)"

run: all
	@ ./$(NAME)
