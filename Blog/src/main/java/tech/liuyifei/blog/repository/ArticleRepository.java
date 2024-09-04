package tech.liuyifei.blog.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import tech.liuyifei.blog.entity.Article;

@Repository
public interface ArticleRepository extends JpaRepository<Article, Long> {
}
