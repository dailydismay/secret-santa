import { Avatar, Navbar } from "@nextui-org/react";
import { Link } from "react-router-dom";
import { Visitor } from "store/visitor/types";

export type ProfileProps = {
  visitor: Visitor;
};

export const Profile: React.FC<ProfileProps> = ({ visitor }) => (
  <Link to={"/profile"}>
    <Avatar
      css={{ cursor: "pointer" }}
      squared
      src={visitor.avatarURL}
      text={visitor.firstName}
    />
  </Link>
);
