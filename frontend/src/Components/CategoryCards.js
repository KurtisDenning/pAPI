import { React } from "react";
import { Button } from "@chakra-ui/react";

const CategoryCards = (props) => {
  function setState() {
    if (props.category === props.name) {
      props.setCategory("");
    } else {
      props.setCategory(props.name);
    }
  }

  return (
    <Button
      m={2}
      isActive={props.category === props.name}
      onClick={() => {
        setState();
      }}
    >
      {props.name}
    </Button>
  );
};

export default CategoryCards;
