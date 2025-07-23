#include <iostream>

// u can make interface with pure virtual functions
// for ex. lets make printable interface
class Printable
{
  virtual std::string get_class_name() = 0;
};

class Entity : public Printable
{
public:
  // this is a virtual function, can be overriden by subclasses
  virtual std::string get_name()
  {
    return "Entity";
  }
  // following is a pure virtual function, has to be overridden by subclasses
  // in context of minecraft, not all entity can float, ex. sand and gravel slides
  virtual bool will_float() = 0;

  std::string get_class_name() override
  {
    return "Entity";
  }
};

class Player : public Entity
{
private:
  std::string name;

public:
  Player(std::string n)
  {
    name = n;
  }

  std::string get_name() override
  {
    return name;
  }

  bool will_float() override
  {
    return false;
  }

  // if u comment out following method,
  // Entity's implementation of the same func will be called
  std::string get_class_name() override
  {
    return "Player";
  }
};

void print_name(Entity *e)
{
  std::cout << e->get_name() << std::endl;
}

void virtual_function_demo()
{
  // u cannot make obj of player unless player overrides pure
  // virtual function
  Player *p = new Player("Hardik");

  std::cout << p->get_class_name() << std::endl;

  print_name(p);
}
