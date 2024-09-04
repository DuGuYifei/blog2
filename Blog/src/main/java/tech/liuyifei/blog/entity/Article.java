package tech.liuyifei.blog.entity;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.OneToMany;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import tech.liuyifei.blog.enums.ArticleStatus;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Entity
public class Article {
    @Id
    private Long id;
    @ManyToOne
    private Article parent;
    @OneToMany(mappedBy = "parent")
    private List<Article> children;
    private ArticleStatus status;
}
