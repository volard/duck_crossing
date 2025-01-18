import React from 'react';
import { observer } from 'mobx-react-lite';
import { Container, Title } from '@mantine/core';
import { useStore } from './store';

const App = observer(() => {
  const { counterStore } = useStore();

  return (
    <Container>
      <Title align="center">Counter: {counterStore.count}</Title>
    </Container>
  );
});

export default App;
