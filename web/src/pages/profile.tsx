import { useStore } from "effector-react";
import { Container, Grid, Col } from "@nextui-org/react";
import { $isLoggedIn, $visitor } from "store/visitor";
import { CreateGroupModal } from "components/profile/create-group-modal";
import { ProfileCard } from "components/profile/profile-card";
import { OwnGroupsCard } from "components/profile/own-groups-card";

export const ProfilePage: React.FC = () => {
  const isLoggedIn = useStore($isLoggedIn);
  const visitor = useStore($visitor);

  if (!isLoggedIn || !visitor) return <></>;

  return (
    <Container css={{ maxW: "1400px", h: "100%", flex: 1 }}>
      <CreateGroupModal />
      <Grid.Container gap={2} justify={"flex-start"}>
        <Grid xl={3} md={3} lg={3} sm={12} xs={12}>
          <Col>
            <ProfileCard visitor={visitor} />
          </Col>
        </Grid>
        <Grid xl={9} md={9} lg={9} sm={12} xs={12}>
          <Col>
            <OwnGroupsCard />
          </Col>
        </Grid>
      </Grid.Container>
    </Container>
  );
};
