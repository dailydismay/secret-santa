import { Navbar, Text } from "@nextui-org/react";
import { Link, useLocation } from "react-router-dom";

export type NavItemLinkProps = {
  path: string;
  currentPath: string;
  text: string;
};

export const NavItemLink: React.FC<NavItemLinkProps> = ({
  path,
  text,
  currentPath,
}) => {
  return (
    <Navbar.Link color={"secondary"} isActive={path === currentPath}>
      <Link to={path} style={{ color: "inherrit" }}>
        <Text color={"default"}>{text}</Text>
      </Link>
    </Navbar.Link>
  );
};
