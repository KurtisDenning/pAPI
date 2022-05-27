import React from "react";
import { Button } from "@chakra-ui/react";

const CategoryCards = (props) => {
  return (
    <Button onClick={() => props.setCategory(props.name)} m={2}>
      {props.name}
    </Button>
  );
};

export default CategoryCards;
