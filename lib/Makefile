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

#####   INFO   #####

NAME			= ../libgopy.so

#####   COMMANDS   #####

CC				= g++
CFLAGS			= -Werror -Wall -Wextra -fPIC -shared -O2 -s
PYFLAGS			= $(shell python3-config --includes) $(shell python3-config --ldflags) -lpython3.10

MD				= mkdir -p
RM				= rm -rf


#####   RESOURCES   #####

SRC_DIR			= srcs
INC_DIR			= includes
OBJ_DIR			= objs

SRC_SUB_DIRS	= $(shell find $(SRC_DIR) -type d)
INC_SUB_DIRS	= $(shell find $(INC_DIR) -type d)
OBJ_SUB_DIRS	= $(SRC_SUB_DIRS:$(SRC_DIR)%=$(OBJ_DIR)%)

SRCS			= $(foreach DIR, $(SRC_SUB_DIRS), $(wildcard $(DIR)/*.cpp))
INCS			= $(foreach DIR, $(INC_SUB_DIRS), $(wildcard $(DIR)/*.h))
OBJS			= $(SRCS:$(SRC_DIR)/%.cpp=$(OBJ_DIR)/%.o)



#####   BASE RULES   #####

all: $(NAME)

$(NAME): $(OBJ_SUB_DIRS) $(OBJS)
	@ $(CC) $(CFLAGS) $(OBJS) -o $@ $(PYFLAGS)
	@ echo "$(GREEN)[+] $(NAME)$(END)"

$(OBJ_DIR)/%.o : $(SRC_DIR)/%.cpp $(INCS)
	@ $(CC) $(CFLAGS) -I$(INC_DIR) -c $< -o $@ $(PYFLAGS)
	@ echo "$(BLUE)[+] $@$(END)"

$(OBJ_SUB_DIRS):
	@ $(MD) $(OBJ_SUB_DIRS)
	@ echo "$(PURPLE)[+] $(SRC_DIR) -> $(OBJ_DIR) $(END)"


clean:
	@ $(RM) $(OBJ_DIR)
	@ echo "$(YELLOW)[+] $(OBJ_DIR)$(END)"

fclean: clean
	@ $(RM) $(NAME)
	@ echo "$(YELLOW)[+] $(NAME)$(END)"

re: fclean all

run:
	@ ./$(NAME)
rrun: fclean run



#####   PHONY   #####

.PHONY: all clean fclean re run rrun


#libgopy.so: libgopy.cpp
#	g++ -shared -o libgopy.so libgopy.cpp $(shell python3-config --includes) $(shell python3-config --ldflags) -lpython3.10 -fPIC
