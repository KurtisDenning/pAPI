import { React } from "react";
import { Button } from "@chakra-ui/react";

const CategoryCards = (props) => {
  function setState() {
    if (props.category === props.shortDesc) {
      props.setCategory("");
    } else {
      props.setCategory(props.shortDesc);
    }
  }

  return (
    <Button
      mr={[6, 9, 12]}
      mb={[6, 9, 12]}
      isActive={props.category === props.shortDesc}
      onClick={() => {
        setState();
      }}
    >
      {props.shortDesc}
    </Button>
  );
};

export default CategoryCards;
